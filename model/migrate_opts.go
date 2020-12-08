package model

import "github.com/yufeifly/proxy/api/types"

type MigrateReqOpts struct {
	Src           types.Address // migration src
	Dst           types.Address // migration destination
	ServiceID     string        // of worker
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}

type MigrateOpts struct {
	types.Address
	ServiceID     string
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}
