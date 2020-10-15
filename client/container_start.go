package client

import (
	"github.com/levigross/grequests"
	"github.com/yufeifly/proxy/model"
)

func (cli *Client) ContainerStart(opts model.StartReqOpts) error {
	data := make(map[string]string)
	data["ContainerID"] = opts.ContainerID
	data["CheckpointID"] = opts.CStartOpts.CheckpointID
	data["CheckpointDir"] = opts.CStartOpts.CheckpointDir

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := "http://" + opts.IP + ":" + opts.Port + "/container/start"
	_, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}

	return nil
}
