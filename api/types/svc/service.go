package svc

import "github.com/yufeifly/proxy/api/types"

type ServiceOpts struct {
	ID             string
	ProxyServiceID string
	NodeAddr       types.Address
}
