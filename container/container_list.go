package container

import (
	dockertypes "github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/client"
)

type ListReqOpts struct {
	types.Address
	dockertypes.ContainerListOptions
}

// ListContainers list containers of a worker node
func ListContainers(listOpts ListReqOpts) ([]dockertypes.Container, error) {
	cli := client.Client{
		Target: listOpts.Address,
	}
	lOpts := types.ListOpts{
		ContainerListOptions: listOpts.ContainerListOptions,
	}
	containers, err := cli.ContainerList(lOpts)
	if err != nil {
		logrus.Errorf("container.ListContainers, cli.ContainerList err: %v", err)
		return nil, err
	}
	return containers, nil
}
