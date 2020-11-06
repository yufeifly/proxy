package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
)

// SendLog send log to dst
func (cli *Client) SendLog(logWithID model.LogWithServiceID) error {
	logrus.Infof("data to send: %v\n", logWithID.Log)
	dataJson, _ := json.Marshal(logWithID)

	ro := &grequests.RequestOptions{
		JSON: dataJson,
	}

	//url := "http://127.0.0.1:6789/logger"
	url := "http://" + cli.Dest.IP + ":" + cli.Dest.Port + "/logger"
	resp, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("client.SendLog post err: %v", err)
		return err
	}
	logrus.Info(resp)
	return nil
}
