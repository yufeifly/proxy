package client

import (
	"github.com/yufeifly/proxy/api/types"
)

// Client ...
type Client struct {
	addr types.Address
}

// NewClient new a client by target address
func NewClient(address types.Address) APIClient {
	return &Client{
		addr: address,
	}
}
