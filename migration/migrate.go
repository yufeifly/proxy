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
func TrySendMigrate(opts model.MigrateReqOpts) error {

	// select a dst node, and open connection to dst

	logrus.Warn("ticket set logging")
	ticket.T.SetTicket(ticket.Logging)

	// send migrate request to src node
	cli := client.Client{}
	err := cli.SendMigrate(opts)
	if err != nil {
		return err
	}

	// write log files to dst
	// when dst starts, open redis connection
	//  dst consume logs in the meantime
	// wait until all log files consumed(no whole log file)
	tick := time.Tick(200 * time.Millisecond)
	for {
		<-tick
		sent := wal.LockAndGetTotalSend()
		consumed := wal.LockAndGetTotal()
		if sent == 0 {
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

	// log, send leftover
	wal.SendLogEntry()

	// wait until the last log consumed by dst
	for {
		<-tick
		sent := wal.LockAndGetTotalSend()
		consumed := wal.LockAndGetTotal()
		if sent == consumed {
			// switch, requests redirect to dst node

			break
		}
	}

	// downtime end, unset global lock
	logrus.Warn("ticket unset")
	ticket.T.UnSet()

	return nil
}
