/*
Q: how to tell the dst consumer goroutine to stop?
A: via function service.SendLastLog()
*/
package migration

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/nodeSelector"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/ticket"
	"github.com/yufeifly/proxy/utils"
	"time"
)

// TryMigrate migrate redis service
func TryMigrateWithLogging(reqOpts model.MigrateReqOpts) error {
	// select an appropriate dst node
	if reqOpts.Dst.IP == "" || reqOpts.Dst.Port == "" {
		node := nodeSelector.BestTarget()
		reqOpts.Dst = node.Address
	}

	// add service.MigrationTarget first, start logging second!
	service, err := scheduler.Default().GetService(reqOpts.ProxyService)
	if err != nil {
		logrus.Errorf("migration.TryMigrateWithLogging GetService failed, err: %v", err)
		return err
	}
	logrus.Debugf("migration.TryMigrateWithLogging service: %v", service)

	addr := model.Address{
		IP:   reqOpts.Dst.IP,
		Port: reqOpts.Dst.Port,
	}
	service.AddMigTarget(addr)

	reqOpts.ServiceID = service.ID // of worker

	ticket.Default().Set(ticket.Logging)

	// for test, test if log are consumed successfully
	// redis.Set("service1", "happy", "birthday")

	started := make(chan bool) // todo change to struct{}{}
	// send migrate request to src node
	go func() {
		cli := client.Client{
			Target: reqOpts.Src,
		}
		err := cli.SendMigrate(reqOpts)
		if err != nil {
			logrus.Panicf("cli.SendMigrate failed, err: %v", err)
		}
		logrus.Debug("container dst started")
		started <- true
		logrus.Debug("container dst started, true write to chan")
	}()

	// write log files to dst
	// when dst starts, open redis connection
	// dst consume logs in the meantime
	// wait until all log files consumed(no whole log file)
	ticker := time.NewTicker(10 * time.Millisecond)
FOR:
	for {
		select {
		case <-started:
			logrus.Debug("get value from chan(started)")
			sent, _ := service.LockAndGetSentConsumed()
			if sent == 0 {
				logrus.Debug("log sent is 0, about to sent last log")
				service.UnlockLogger()
				break FOR
			}
			service.UnlockLogger()
		case <-ticker.C:
			//logrus.Info("tick")
			sent, consumed := service.LockAndGetSentConsumed()
			//logrus.Infof("sent: %v, consumed: %v", sent, consumed)
			if sent == 0 {
				service.UnlockLogger()
				continue
			}
			if sent-consumed < 1 {
				logrus.Warn("downtime start")
				ticket.Default().Set(ticket.ShutWrite)
				service.UnlockLogger()
				break FOR
			}
			service.UnlockLogger()
		}
	}

	// send the last log with flag "true" to dst,
	// true flag tells dst that this is the last one, so the consumer goroutine can stop
	err = service.SendLastLog()
	if err != nil {
		logrus.Errorf("migration service.SendLastLog failed, err: %v", err)
		return err
	}

	// wait until the last log consumed by dst
	for {
		<-ticker.C
		sent, consumed := service.LockAndGetSentConsumed()
		if sent == consumed {
			service.UnlockLogger()
			logrus.Debug("switching, requests redirect to dst node")
			opts := model.ServiceOpts{
				ID:             utils.RenameService(reqOpts.ServiceID),
				ProxyServiceID: reqOpts.ProxyService,
				NodeAddr: model.Address{
					IP:   reqOpts.Dst.IP,
					Port: reqOpts.Dst.Port,
				},
			}
			scheduler.DefaultRegister(reqOpts.ProxyService, opts)
			logrus.Warn("downtime end")
			break
		}
		service.UnlockLogger()
	}
	ticker.Stop() // shut ticker

	// downtime end, unset global lock
	logrus.Debug("ticket unset")
	ticket.Default().UnSet()

	return nil
}

// TryMigrate migrate regular containers
func TryMigrate(reqOpts model.MigrateReqOpts) error {
	// select an appropriate dst node
	if reqOpts.Dst.IP == "" || reqOpts.Dst.Port == "" {
		node := nodeSelector.BestTarget()
		reqOpts.Dst = node.Address
	}

	// send migrate request to src node

	cli := client.Client{
		Target: reqOpts.Src,
	}
	err := cli.SendMigrate(reqOpts)
	if err != nil {
		logrus.Panicf("cli.SendMigrate failed, err: %v", err)
	}

	// downtime end, unset global lock
	logrus.Debug("ticket unset")
	ticket.Default().UnSet()

	return nil
}
