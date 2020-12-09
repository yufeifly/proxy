package svc

import "github.com/yufeifly/proxy/api/types"

// ServiceOpts ...
type ServiceOpts struct {
	ID             string
	ProxyServiceID string
	NodeAddr       types.Address
}
