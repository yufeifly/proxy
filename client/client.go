package client

import (
	"fmt"
	"github.com/yufeifly/proxy/api/types"
)

// Client ...
type Client struct {
	addr types.Address
}

func NewClient(address types.Address) APIClient {
	return &Client{
		addr: address,
	}
}

// getAPIPath path means webapi path, for example: /redis/set
func (cli *Client) getAPIPath(path string) string {
	return fmt.Sprintf("http://%s:%s%s", cli.addr.IP, cli.addr.Port, path)
}
