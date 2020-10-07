package migration

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/ticket"
	"time"
)

// TryMigrate migrate redis service
func TryMigrate(opts model.MigrateOpts) error {

	// select a dst node, and open connection to dst

	// set global lock
	ticket.T.Set()
	logrus.Warn("ticket set")

	// write log files to dst

	// send migrate request to src node

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
	ticket.T.UnSet()
	logrus.Warn("ticket unset")
	return nil
}
