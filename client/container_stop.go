package client

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
)

// StopContainer ...
func (cli *Client) StopContainer(options types.StopOpts) error {
	data := make(map[string]string, 2)
	data["ContainerID"] = options.ContainerID
	data["Timeout"] = options.Timeout

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
