package migration

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/ticket"
	"time"
)

// TryMigrate migrate redis service
func TrySendMigrate(opts model.MigrateReqOpts) error {

	// select a dst node, and open connection to dst

	// set global lock
	ticket.T.SetTicket(ticket.Logging)
	logrus.Warn("ticket set logging")

	// write log files to dst

	// send migrate request to src node
	cli := client.Client{}
	err := cli.SendMigrate(opts)
	if err != nil {
		return err
	}

	time.Sleep(60 * time.Second)

	// when dst starts, open redis connection

	//  dst consume logs in the meantime

	// wait until all log files consumed(no whole log file)

	// downtime start

	// send last log file to dst to let it consume

	// dst no log file

	// switch, requests redirect to dst node

	// downtime end
	// unset global lock
	logrus.Warn("ticket unset")
	return nil
}
