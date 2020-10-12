package container

import (
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
)

// StartContainer start a container with opts
func StartContainer(opts model.StartOpts) error {
	cli := client.Client{}
	err := cli.ContainerStart(opts)
	if err != nil {
		return nil
	}
	return nil
}
