package client

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/scheduler"
)

func (cli *Client) AddService(service *scheduler.Service) error {
	data := make(map[string]string)
	data["ServiceID"] = service.ID
	data["ProxyServiceID"] = service.ProxyServiceID

	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := "http://" + service.Node.IP + ":" + service.Node.Port + "/service/add"
	resp, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("AddService.post err : %v", err)
		return err
	}
	logrus.Infof("AddService resp: %v", resp.RawResponse)
	return nil
}
