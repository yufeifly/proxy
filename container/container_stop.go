package container

import (
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
)

func StopContainer(opts model.StopReqOpts) error {
	cli := client.Client{}
	err := cli.StopContainer(opts)
	if err != nil {
		return err
	}
	return nil
}
