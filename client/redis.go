package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/yufeifly/proxy/model"
)

// RedisGet send get request to worker node
func (cli *Client) RedisGet(opts model.RedisGetOpts) (string, error) {
	params := make(map[string]string, 2)
	params["key"] = opts.Key
	params["service"] = opts.ServiceID
	ro := &grequests.RequestOptions{
		Params: params,
	}
	url := "http://" + opts.Node.IP + ":" + opts.Node.Port + "/redis/get"
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
func (cli *Client) RedisSet(opts model.RedisSetOpts) error {
	data := make(map[string]string, 3)
	data["key"] = opts.Key
	data["value"] = opts.Value
	data["service"] = opts.ServiceID

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := "http://" + opts.Node.IP + ":" + opts.Node.Port + "/redis/set"
	_, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}
	return nil
}
