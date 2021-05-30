/*
Package migration
Q: how to tell the dst consumer goroutine to stop?
A: via function service.SendLastLog()
*/
package migration

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/api/types/svc"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/nodeSelector"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/ticket"
	"github.com/yufeifly/proxy/utils"
	"time"
)

// MigrateReqOpts ...
type MigrateReqOpts struct {
	Src           types.Address // migration src
	Dst           types.Address // migration destination
	ServiceID     string        // of worker
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}

// TryMigrateWithLogging migrate redis service
func TryMigrateWithLogging(reqOpts MigrateReqOpts) error {
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

	addr := types.Address{
		IP:   reqOpts.Dst.IP,
		Port: reqOpts.Dst.Port,
	}
	service.AddMigTarget(addr)

	reqOpts.ServiceID = service.ID // of worker

	service.Ticket().Set(ticket.Logging)

	startedCh := make(chan bool) // todo change to struct{}{}
	// send migrate request to src node
	go func() {
		cli := client.NewClient(reqOpts.Src)
		mOpts := types.MigrateOpts{
			Address:       reqOpts.Dst,
			ServiceID:     reqOpts.ServiceID,
			ProxyService:  reqOpts.ProxyService,
			CheckpointID:  reqOpts.CheckpointID,
			CheckpointDir: reqOpts.CheckpointDir,
		}
		err := cli.SendMigrate(mOpts)
		if err != nil {
			logrus.Panicf("cli.SendMigrate failed, err: %v", err)
		}
		logrus.Debug("container dst started")
		startedCh <- true
		logrus.Debug("container dst started, true write to chan")
	}()

	// write log files to dst
	// when dst starts, open redis connection
	// dst consume logs in the meantime
	// wait until all log files are consumed(no whole log file)
	ticker := time.NewTicker(1 * time.Microsecond)
FOR:
	for {
		select {
		case <-startedCh:
			logrus.Debug("migration.TryMigrateWithLogging, get value from chan(started)")
			sent, _ := service.LockAndGetSentConsumed()
			if sent == 0 {
				logrus.Debug("migration.TryMigrateWithLogging, log sent number is 0, about to send the last log")
				break FOR
			}
		case <-ticker.C:
			sent, consumed := service.LockAndGetSentConsumed()
			if sent == 0 {
				continue
			}
			if sent-consumed < 1 {
				logrus.Warn("migration.TryMigrateWithLogging, downtime start")
				service.Ticket().Set(ticket.ShutWrite)
				break FOR
			}
		}
	}

	// send the last log with flag "true" to dst,
	// true flag tells dst that this is the last one, so the consumer goroutine can stop
	err = service.SendLastLog()
	if err != nil {
		logrus.Errorf("migration.TryMigrateWithLogging, SendLastLog failed, err: %v", err)
		return err
	}

	// wait until the last log consumed by dst
	for {
		<-ticker.C
		sent, consumed := service.LockAndGetSentConsumed()
		if sent == consumed {
			logrus.Warn("migration.TryMigrateWithLogging, switching, requests redirect to dst node")
			opts := svc.ServiceOpts{
				ID:             utils.RenameService(reqOpts.ServiceID),
				ProxyServiceID: reqOpts.ProxyService,
				NodeAddr: types.Address{
					IP:   reqOpts.Dst.IP,
					Port: reqOpts.Dst.Port,
				},
			}
			scheduler.DefaultRegister(reqOpts.ProxyService, opts)
			logrus.Warn("migration.TryMigrateWithLogging, downtime end")
			break
		}
	}
	ticker.Stop() // shut ticker

	// downtime end, unset global lock
	logrus.Debug("migration.TryMigrateWithLogging, ticket unset")
	service.Ticket().UnSet()

	return nil
}

// TryMigrate migrate regular containers
func TryMigrate(reqOpts MigrateReqOpts) error {
	// select an appropriate dst node
	if reqOpts.Dst.IP == "" || reqOpts.Dst.Port == "" {
		node := nodeSelector.BestTarget()
		reqOpts.Dst = node.Address
	}

	// add service.MigrationTarget first, start logging second!
	service, err := scheduler.Default().GetService(reqOpts.ProxyService)
	if err != nil {
		logrus.Errorf("migration.TryMigrate GetService failed, err: %v", err)
		return err
	}
	logrus.Debugf("migration.TryMigrate service: %v", service)

	reqOpts.ServiceID = service.ID // of worker

	// send migrate request to src node
	cli := client.NewClient(reqOpts.Src)
	mOpts := types.MigrateOpts{
		Address:       reqOpts.Dst,
		ServiceID:     reqOpts.ServiceID,
		ProxyService:  reqOpts.ProxyService,
		CheckpointID:  reqOpts.CheckpointID,
		CheckpointDir: reqOpts.CheckpointDir,
	}
	if err := cli.SendMigrate(mOpts); err != nil {
		logrus.Panicf("migration.TryMigrate, cli.SendMigrate failed, err: %v", err)
	}

	logrus.Warn("migration.TryMigrateWithLogging, switching, requests redirect to dst node")
	opts := svc.ServiceOpts{
		ID:             utils.RenameService(reqOpts.ServiceID),
		ProxyServiceID: reqOpts.ProxyService,
		NodeAddr: types.Address{
			IP:   reqOpts.Dst.IP,
			Port: reqOpts.Dst.Port,
		},
	}
	scheduler.DefaultRegister(reqOpts.ProxyService, opts)
	logrus.Warn("migration.TryMigrateWithLogging, downtime end")

	logrus.Debug("ticket unset")
	service.Ticket().UnSet()

	return nil
}
