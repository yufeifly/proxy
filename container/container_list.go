package container

import (
	"github.com/docker/docker/api/types"
	"github.com/yufeifly/proxy/client"
)

func ListContainers(opts types.ContainerListOptions) ([]types.Container, error) {
	cli := client.Client{}
	containers, err := cli.ContainerList(opts)
	if err != nil {
		return nil, err
	}
	return containers, nil
}
