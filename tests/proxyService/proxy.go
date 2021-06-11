package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
)

func TestProxyService() error {
	ro := &grequests.RequestOptions{
		Params: map[string]string{"Service": "s1.c1"},
	}

	url := fmt.Sprintf("http://%s:%s/proxy/service/get", "192.168.134.138", "6788")
	resp, err := grequests.Get(url, ro)
	if err != nil {
		return err
	}

	var node types.Address
	json.Unmarshal(resp.Bytes(), &node)
	logrus.Infof("ip: %v, port: %v", node.IP, node.Port)

	resp.RawResponse.Body.Close()

	return nil
}

func main() {
	err := TestProxyService()
	if err != nil {
		logrus.Error(err)
	}
}
