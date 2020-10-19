package client

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/model"
)

// SendLog send log to dst
func (cli *Client) SendLog(logWithID model.LogWithServiceID) error {
	fmt.Printf("data to send: %v\n", logWithID.Log)
	dataJson, _ := json.Marshal(logWithID)
	//dataPost := make(map[string]string)
	//dataPost["Service"] = logWithID.ProxyServiceID

	ro := &grequests.RequestOptions{
		JSON: dataJson,
		//Data: dataPost,
	}

	//url := "http://127.0.0.1:6789/logger"
	url := "http://" + cli.Dest.IP + ":" + cli.Dest.Port + "/logger"
	resp, err := grequests.Post(url, ro)
	if err != nil {
		return err
	}
	logrus.Info(resp)
	return nil
}
