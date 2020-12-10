package client

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types/logger"
)

// SendLog send log to dst
func (cli *Client) SendLog(logWithID logger.LogWithServiceID) error {
	logrus.Debugf("data to send: %v", logWithID.Log)
	dataJSON, err := json.Marshal(logWithID)
	if err != nil {
		logrus.Errorf("client.SendLog Marshal failed, err :%v", err)
		return err
	}

	ro := &grequests.RequestOptions{
		JSON: dataJSON,
	}

	//url := "http://127.0.0.1:6789/logger"
	//url := "http://" + cli.Target.IP + ":" + cli.Target.Port + "/logger"
	url := cli.getAPIPath("/logger")
	resp, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("client.SendLog Post failed, err: %v", err)
		return err
	}
	logrus.Info(resp)
	return nil
}