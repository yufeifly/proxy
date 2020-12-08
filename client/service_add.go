package client

import (
	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types/svc"
)

// AddService ...
func (cli *Client) AddService(service svc.ServiceOpts) error {
	data := make(map[string]string, 3)
	data["ServiceID"] = service.ID
	data["ProxyServiceID"] = service.ProxyServiceID
	ro := &grequests.RequestOptions{
		Data: data,
	}
	url := cli.getAPIPath("/service/add")
	resp, err := grequests.Post(url, ro)
	if err != nil {
		logrus.Errorf("AddService.post err : %v", err)
		return err
	}
	logrus.Infof("AddService resp: %v", resp.RawResponse)
	return nil
}
