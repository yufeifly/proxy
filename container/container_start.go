package container

import (
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
)

// StartContainer start a container with opts
func StartContainer(opts model.StartReqOpts) error {
	cli := client.Client{
		Target: opts.Address,
	}
	err := cli.ContainerStart(opts)
	if err != nil {
		return err
	}
	return nil
}
