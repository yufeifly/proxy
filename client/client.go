package client

import (
	"github.com/yufeifly/proxy/api/types"
	"net/http"
	"time"
)

var defaultClient APIClient

func init() {
	defaultClient = NewDefaultClient()
}

// Client ...
type Client struct {
	addr       types.Address
	httpClient *http.Client
}

// NewClient new a client by target address
func NewClient(address types.Address) APIClient {
	return &Client{
		addr: address,
	}
}

func NewDefaultClient() APIClient {
	return &Client{
		httpClient: createHTTPClient(),
	}
}

func DefaultClient() APIClient {
	return defaultClient
}

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}
