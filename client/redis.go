package client

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"net/http"
	"net/url"
)

// RedisGet send get request to worker node
func (cli *Client) RedisGet(options types.RedisGetOpts) (string, error) {
	logrus.Infof("target: %v", options.Node)
	url := fmt.Sprintf("http://%s:%s%s", options.Node.IP, options.Node.Port, "/redis/get")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logrus.Errorf("http.NewRequest err: %v", err)
		return "", err
	}
	// add params
	q := req.URL.Query()
	q.Add("key", options.Key)
	q.Add("service", options.ServiceID)
	req.URL.RawQuery = q.Encode()
	// do request
	resp, err := cli.httpClient.Do(req)
	if err != nil {
		logrus.Errorf("cli.httpClient.Do err: %v", err)
		return "", err
	}
	defer resp.Body.Close()
	var value string
	err = json.NewDecoder(resp.Body).Decode(&value)
	if err != nil {
		logrus.Errorf("client.RedisGet Read json.Unmarshal err: %v", err)
		return "", err
	}
	return value, nil
}

// RedisSet send set request to worker node
func (cli *Client) RedisSet(options types.RedisSetOpts) error {
	urlSet := fmt.Sprintf("http://%s:%s%s", options.Node.IP, options.Node.Port, "/redis/set")
	// add params
	q := url.Values{
		"key":     {options.Key},
		"value":   {options.Value},
		"service": {options.ServiceID},
	}
	// do request
	resp, err := cli.httpClient.PostForm(urlSet, q)
	if err != nil {
		logrus.Errorf("cli.httpClient.Do err: %v", err)
		return err
	}
	resp.Body.Close()
	return nil
}

func (cli *Client) RedisDelete(options types.RedisDeleteOpts) error {
	data := make(map[string]string, 2)
	data["key"] = options.Key
	data["service"] = options.ServiceID

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := cli.getAPIPath("/redis/delete")
	resp, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}
	resp.RawResponse.Body.Close()
	return nil
}
