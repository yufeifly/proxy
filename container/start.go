package container

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
)

// StartContainer start a container with opts
func StartContainer(startOpts model.StartOpts) error {
	header := "container.StartContainer"
	err := cli.ContainerStart(context.Background(), startOpts.ContainerID, startOpts.CStartOpts)
	if err != nil {
		logrus.Errorf("%s, start container failed, err: %v", header, err)
		return err
	}
	return nil
}
