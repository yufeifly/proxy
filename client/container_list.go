package client

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/levigross/grequests"
	"github.com/yufeifly/proxy/model"
)

// ContainerList send request to target node to get the containers's info
func (cli *Client) ContainerList(opts model.ListReqOpts) ([]types.Container, error) {
	listOptsJson, _ := json.Marshal(opts.ContainerListOptions)

	ro := &grequests.RequestOptions{
		JSON: listOptsJson,
	}
	url := "http://" + opts.IP + ":" + opts.Port + "/container/list"
	resp, err := grequests.Get(url, ro)
	if err != nil {
		return nil, err
	}

	var containers []types.Container
	err = json.NewDecoder(resp.RawResponse.Body).Decode(&containers)
	ensureReaderClosed(resp)

	return containers, nil
}
