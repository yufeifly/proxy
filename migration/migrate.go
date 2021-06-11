/*
Package migration
Q: how to tell the dst consumer goroutine to stop?
A: via function service.SendLastLog()
*/
package migration

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/nodeSelector"
	"github.com/yufeifly/proxy/scheduler"
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
func MigrateWithLogging(reqOpts MigrateReqOpts) error {
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

	// send migrate request to src node

	cli := client.NewClient(reqOpts.Src)
	mOpts := types.MigrateOpts{
		Address:       reqOpts.Dst,
		ServiceID:     reqOpts.ServiceID,
		ProxyService:  reqOpts.ProxyService,
		CheckpointID:  reqOpts.CheckpointID,
		CheckpointDir: reqOpts.CheckpointDir,
	}
	err = cli.SendMigrate(mOpts)
	if err != nil {
		logrus.Panicf("cli.SendMigrate failed, err: %v", err)
	}
	logrus.Debug("container dst started")
	logrus.Debug("container dst started, true write to chan")

	return nil
}
