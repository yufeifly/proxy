package client

import (
	"fmt"
	"github.com/yufeifly/proxy/model"
)

// Client ...
type Client struct {
	Target model.Address
}

// getAPIPath path means webapi path, for example: /redis/set
func (cli *Client) getAPIPath(path string) string {
	return fmt.Sprintf("http://%s:%s%s", cli.Target.IP, cli.Target.Port, path)
}
