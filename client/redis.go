package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/yufeifly/proxy/api/types"
)

// RedisGet send get request to worker node
func (cli *Client) RedisGet(options types.RedisGetOpts) (string, error) {
	params := make(map[string]string, 2)
	params["key"] = options.Key
	params["service"] = options.ServiceID
	ro := &grequests.RequestOptions{
		Params: params,
	}
	url := cli.getAPIPath("/redis/get")
	resp, err := grequests.Get(url, ro)
	if err != nil {
		return "", err
	}
	var value string
	err = json.NewDecoder(resp.RawResponse.Body).Decode(&value)
	ensureReaderClosed(resp)
	return value, nil
}

// RedisSet send set request to worker node
func (cli *Client) RedisSet(options types.RedisSetOpts) error {
	data := make(map[string]string, 3)
	data["key"] = options.Key
	data["value"] = options.Value
	data["service"] = options.ServiceID

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := cli.getAPIPath("/redis/set")
	_, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}
	return nil
}
