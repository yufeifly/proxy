package client

import (
	"encoding/json"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
)

// ContainerList send request to target node to get the containers's info
func (cli *Client) ContainerList(opts types.ListOpts) ([]dockertypes.Container, error) {
	listOptsJSON, err := json.Marshal(opts.ContainerListOptions)
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
