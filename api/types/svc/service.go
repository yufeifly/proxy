package svc

import "github.com/yufeifly/proxy/api/types"

// ServiceOpts ...
type ServiceOpts struct {
	CName    string
	SID      string
	NodeAddr types.Address
}
