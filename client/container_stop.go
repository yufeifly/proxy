package client

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
)

func (cli *Client) StopContainer(opts model.StopReqOpts) error {
	data := make(map[string]string, 2)
	data["ContainerID"] = opts.ContainerID
	data["Timeout"] = opts.Timeout

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := cli.getAPIPath("/container/stop")
	_, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("client.StopContainer post err: %v", err)
		return err
	}
	return nil
}
