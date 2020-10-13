package client

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
)

// SendLog send log to dst
func (cli *Client) SendLog(data model.Log) error {
	fmt.Printf("data to send: %v\n", data)
	dataJson, _ := json.Marshal(data)
	ro := &grequests.RequestOptions{
		JSON: dataJson,
	}

	url := "http://127.0.0.1:6789/logger"
	resp, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}
	logrus.Info(resp)
	return nil
}
