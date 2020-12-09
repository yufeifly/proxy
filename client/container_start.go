package client

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
)

// ContainerStart send start request
func (cli *Client) ContainerStart(options types.StartOpts) error {
	data := make(map[string]string, 3)
	data["ContainerID"] = options.ContainerID
	data["CheckpointID"] = options.CStartOpts.CheckpointID
	data["CheckpointDir"] = options.CStartOpts.CheckpointDir

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := cli.getAPIPath("/container/start")
	_, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("client.ContainerStart post err: %v", err)
		return err
	}
	return nil
}
