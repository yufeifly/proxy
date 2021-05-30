package client

import (
	"encoding/json"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
)

// ContainerList send request to target node to get the containers's info
func (cli *Client) ContainerList(options types.ListOpts) ([]dockertypes.Container, error) {
	listOptsJSON, err := json.Marshal(options.ContainerListOptions)
	if err != nil {
		logrus.Errorf("client.ContainerList Marshal failed, err : %v", err)
		return nil, err
	}

	ro := &grequests.RequestOptions{
		JSON: listOptsJSON,
	}
	url := cli.getAPIPath("/container/list")
	resp, err := grequests.Get(url, ro)
	if err != nil {
		return nil, err
	}

	var containers []dockertypes.Container
	err = json.NewDecoder(resp.RawResponse.Body).Decode(&containers)
	ensureReaderClosed(resp)

	return containers, nil
}

// StopContainer ...
func (cli *Client) StopContainer(options types.StopOpts) error {
	data := make(map[string]string, 2)
	data["ContainerID"] = options.ContainerID
	data["Timeout"] = options.Timeout

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := cli.getAPIPath("/container/stop")
	resp, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("client.StopContainer post err: %v", err)
		return err
	}
	resp.RawResponse.Body.Close()
	return nil
}

// ContainerStart send request for starting a container
func (cli *Client) ContainerStart(options types.StartOpts) error {
	data := make(map[string]string, 3)
	data["ContainerID"] = options.ContainerID
	data["CheckpointID"] = options.CStartOpts.CheckpointID
	data["CheckpointDir"] = options.CStartOpts.CheckpointDir

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := cli.getAPIPath("/container/start")
	resp, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("client.ContainerStart post err: %v", err)
		return err
	}
	resp.RawResponse.Body.Close()
	return nil
}
