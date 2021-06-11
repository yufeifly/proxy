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
	SID           string
	CName         string
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
	// verify that service dose exist
	service, err := scheduler.Default().GetService(reqOpts.SID)
	if err != nil {
		logrus.Errorf("migration.TryMigrateWithLogging GetService failed, err: %v", err)
		return err
	}

	// send migrate request to src node
	cli := client.NewClient(reqOpts.Src)
	mOpts := types.MigrateOpts{
		Address:       reqOpts.Dst,
		SID:           reqOpts.SID,
		CName:         reqOpts.CName,
		CheckpointID:  reqOpts.CheckpointID,
		CheckpointDir: reqOpts.CheckpointDir,
	}

	err = cli.SendMigrate(mOpts)
	if err != nil {
		logrus.Panicf("cli.SendMigrate failed, err: %v", err)
	}

	service.UpdateNode(reqOpts.Dst)

	return nil
}
