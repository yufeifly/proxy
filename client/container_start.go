package client

import (
	"github.com/levigross/grequests"
	"github.com/yufeifly/proxy/model"
)

func (cli *Client) ContainerStart(opts model.StartOpts) error {
	data := make(map[string]string)
	data["ContainerID"] = opts.ContainerID
	data["CheckpointID"] = opts.CStartOpts.CheckpointID
	data["CheckpointDir"] = opts.CStartOpts.CheckpointDir

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := "http://127.0.0.1:6789/container/start"
	_, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}

	return nil
}
