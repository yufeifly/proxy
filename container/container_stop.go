package container

import (
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/client"
)

type StopReqOpts struct {
	types.Address
	types.StopOpts
}

// StopContainer stop container
func StopContainer(opts StopReqOpts) error {
	cli := client.NewClient(opts.Address)
	stopOpts := types.StopOpts{
		ContainerID: opts.ContainerID,
		Timeout:     opts.Timeout,
	}
	err := cli.StopContainer(stopOpts)
	if err != nil {
		return err
	}
	return nil
}
