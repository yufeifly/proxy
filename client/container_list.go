package client

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/levigross/grequests"
)

func (cli *Client) ContainerList(opts types.ContainerListOptions) ([]types.Container, error) {
	ro := &grequests.RequestOptions{
		QueryStruct: opts,
	}
	url := "http://127.0.0.1:6789/container/list"
	resp, err := grequests.Get(url, ro)
	if err != nil {
		return nil, err
	}

	var containers []types.Container
	err = json.NewDecoder(resp.RawResponse.Body).Decode(&containers)
	ensureReaderClosed(resp)

	return containers, nil
}
