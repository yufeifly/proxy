package client

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
)

// ContainerList send request to target node to get the containers's info
func (cli *Client) ContainerList(opts model.ListReqOpts) ([]types.Container, error) {
	listOptsJson, err := json.Marshal(opts.ContainerListOptions)
	if err != nil {
		logrus.Errorf("client.ContainerList Marshal failed, err : %v", err)
		return nil, err
	}

	ro := &grequests.RequestOptions{
		JSON: listOptsJson,
	}
	url := cli.getAPIPath("/container/list")
	resp, err := grequests.Get(url, ro)
	if err != nil {
		return nil, err
	}

	var containers []types.Container
	err = json.NewDecoder(resp.RawResponse.Body).Decode(&containers)
	ensureReaderClosed(resp)

	return containers, nil
}
