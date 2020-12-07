package model

import "github.com/docker/docker/api/types"

// StartOpts ...
type StartOpts struct {
	CStartOpts  types.ContainerStartOptions
	ContainerID string
}

// StartReqOpts ...
type StartReqOpts struct {
	Address
	StartOpts
}
