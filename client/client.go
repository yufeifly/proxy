package client

import (
	"fmt"
	"github.com/yufeifly/proxy/model"
)

type Client struct {
	Target model.Address
}

func (cli *Client) getAPIPath(ip, port, path string) string {
	return fmt.Sprintf("http://%s:%s%s", ip, port, path)
}
