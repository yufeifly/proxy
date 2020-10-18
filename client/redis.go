package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/yufeifly/proxy/scheduler"
)

// RedisGet send get request to worker node
func (cli *Client) RedisGet(service *scheduler.Service, key string) (string, error) {
	params := make(map[string]string, 2)
	params["key"] = key
	params["service"] = service.ID
	ro := &grequests.RequestOptions{
		Params: params,
	}
	url := "http://" + service.Node.IP + ":" + service.Node.Port + "/redis/get"
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
func (cli *Client) RedisSet(service *scheduler.Service, key, val string) error {
	data := make(map[string]string, 3)
	data["key"] = key
	data["value"] = val
	data["service"] = service.ID

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := "http://" + service.Node.IP + ":" + service.Node.Port + "/redis/set"
	_, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}
	return nil
}
