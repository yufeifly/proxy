package container

import (
	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
)

// ListContainers list containers of a worker node
func ListContainers(opts model.ListReqOpts) ([]types.Container, error) {
	header := "container.ListContainers"
	cli := client.Client{
		Target: opts.Address,
	}
	containers, err := cli.ContainerList(opts)
	if err != nil {
		logrus.Errorf("%s, cli.ContainerList err: %v", header, err)
		return nil, err
	}
	return containers, nil
}
