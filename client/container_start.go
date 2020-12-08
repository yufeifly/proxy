package client

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
)

// ContainerStart send start request
func (cli *Client) ContainerStart(opts types.StartOpts) error {
	data := make(map[string]string, 3)
	data["ContainerID"] = opts.ContainerID
	data["CheckpointID"] = opts.CStartOpts.CheckpointID
	data["CheckpointDir"] = opts.CStartOpts.CheckpointDir

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
