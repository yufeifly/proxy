/*
Q: how to tell the dst consumer goroutine to stop?
A: via function wal.SendLastLog()
*/
package migration

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/ticket"
	"github.com/yufeifly/proxy/utils"
	"github.com/yufeifly/proxy/wal"
	"time"
)

// TryMigrate migrate redis service
func TrySendMigrate(reqOpts model.MigrateReqOpts) error {
	logrus.Infof("TrySendMigrate.reqOpts: %v", reqOpts)
	// todo select a dst node
	if reqOpts.Dst.IP == "" || reqOpts.Dst.Port == "" {
		reqOpts.Src.IP = "127.0.0.1"
		reqOpts.Src.Port = "6789"
	}

	// add service.Shadow first, start logging second!
	service, _ := scheduler.DefaultScheduler.GetService(reqOpts.ProxyService)
	addr := model.Address{
		IP:   reqOpts.Dst.IP,
		Port: reqOpts.Dst.Port,
	}
	logrus.Infof("TrySendMigrate.service: %v", service)
	if service == nil {
		return nil
	}
	service.AddShadow(addr)
	reqOpts.ServiceID = service.ID // of worker

	logrus.Warn("ticket set logging")
	ticket.T.SetTicket(ticket.Logging)

	started := make(chan bool)
	// send migrate request to src node
	go func() error {
		cli := client.Client{}
		err := cli.SendMigrate(reqOpts)
		if err != nil {
			logrus.Errorf("cli.SendMigrate failed, err: %v", err)
			return err
		}
		logrus.Warn("container dst started")
		started <- true
		logrus.Warn("container dst started, true write to chan")
		return nil
	}()

	// write log files to dst
	// when dst starts, open redis connection
	// dst consume logs in the meantime
	// wait until all log files consumed(no whole log file)
	ticker := time.NewTicker(500 * time.Millisecond)
FOR:
	for {
		select {
		case <-started:
			logrus.Warn("get value from chan(started)")
			sent, _ := wal.LockAndGetSentConsumed()
			if sent == 0 {
				wal.UnlockLogger()
				break FOR
			}
		case <-ticker.C:
			logrus.Info("ticker")
			sent, consumed := wal.LockAndGetSentConsumed()
			logrus.Infof("sent: %v, consumed: %v", sent, consumed)
			if sent == 0 {
				wal.UnlockLogger()
				continue
			}
			if sent-consumed < 1 {
				logrus.Warn("downtime start")
				ticket.T.SetTicket(ticket.ShutWrite)
				wal.UnlockLogger()
				break FOR
			}
			wal.UnlockLogger()
		}
	}

	// send the last log with flag "true" to dst,
	// true flag tells dst that this is the last one, so the consumer goroutine can stop
	err := wal.SendLastLog(reqOpts.ProxyService, addr)
	if err != nil {
		logrus.Errorf("wal.SendLastLog failed, err: %v", err)
		return err
	}

	// wait until the last log consumed by dst
	for {
		<-ticker.C
		sent, consumed := wal.LockAndGetSentConsumed()
		if sent == consumed {
			wal.UnlockLogger()
			// todo switch, requests redirect to dst node
			logrus.Info("switching, requests redirect to dst node")
			opts := model.ServiceOpts{
				ID:             utils.MakeNameForService(reqOpts.ServiceID),
				ProxyServiceID: reqOpts.ProxyService,
				NodeAddr: model.Address{
					IP:   reqOpts.Dst.IP,
					Port: reqOpts.Dst.Port,
				},
			}
			logrus.Infof("new ServiceOpts: %v", opts)
			scheduler.Register(reqOpts.ProxyService, opts)
			logrus.Info("downtime end")
			break
		}
		wal.UnlockLogger()
	}
	ticker.Stop() // shut ticker

	// downtime end, unset global lock
	logrus.Warn("ticket unset")
	ticket.T.UnSet()

	return nil
}
