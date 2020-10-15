package client

import (
	"github.com/levigross/grequests"
	"github.com/yufeifly/proxy/model"
)

func (cli *Client) StopContainer(opts model.StopReqOpts) error {
	data := make(map[string]string)
	data["ContainerID"] = opts.ContainerID
	data["Timeout"] = opts.Timeout

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := "http://" + opts.IP + ":" + opts.Port + "/container/stop"
	_, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}
}
