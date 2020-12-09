package container

import (
	dockertypes "github.com/docker/docker/api/types"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/client"
)

// StartReqOpts ...
type StartReqOpts struct {
	types.Address
	types.StartOpts
}

// StartContainer start a container with opts
func StartContainer(opts StartReqOpts) error {
	cli := client.NewClient(opts.Address)
	sOpts := types.StartOpts{
		CStartOpts:  dockertypes.ContainerStartOptions{},
		ContainerID: "",
	}
	err := cli.ContainerStart(sOpts)
	if err != nil {
		return err
	}
	return nil
}
