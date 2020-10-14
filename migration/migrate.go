package migration

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/ticket"
	"github.com/yufeifly/proxy/wal"
	"time"
)

// TryMigrate migrate redis service
func TrySendMigrate(reqOpts model.MigrateReqOpts) error {

	if reqOpts.Dst.IP == "" || reqOpts.Dst.Port == "" {
		// todo select a dst node, and open connection to dst
	}

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
		return nil
	}()

	// write log files to dst
	// when dst starts, open redis connection
	//  dst consume logs in the meantime
	// wait until all log files consumed(no whole log file)
	ticker := time.NewTicker(1000 * time.Millisecond)
FOR:
	for {
		select {
		case <-started:
			logrus.Warn("get value from chan(started)")
			if wal.LockAndGetTotalSend() == 0 {
				wal.UnlockLogger()
				break FOR
			}
		case <-ticker.C:
			logrus.Info("tick")
			sent := wal.LockAndGetTotalSend()
			consumed := wal.LockAndGetTotalConsumed()
			if sent == 0 {
				wal.UnlockConsumer()
				wal.UnlockLogger()
				continue
			}
			if sent-consumed < 1 {
				logrus.Warn("downtime start")
				ticket.T.SetTicket(ticket.ShutWrite)
				wal.UnlockConsumer()
				wal.UnlockLogger()
				break
			}
			wal.UnlockConsumer()
			wal.UnlockLogger()
		}
	}

	// end leftover log
	err := wal.SendLastLog()
	if err != nil {
		logrus.Errorf("wal.SendLastLog failed, err: %v", err)
		return err
	}

	// wait until the last log consumed by dst
	for {
		<-ticker.C
		sent := wal.LockAndGetTotalSend()
		consumed := wal.LockAndGetTotalConsumed()
		if sent == consumed {
			// switch, requests redirect to dst node
			logrus.Info("switch, requests redirect to dst node")
			logrus.Info("downtime end")
			break
		}
	}
	ticker.Stop() // shut ticker

	// downtime end, unset global lock
	logrus.Warn("ticket unset")
	ticket.T.UnSet()

	// tell the dst consumer goroutine to stop

	return nil
}
