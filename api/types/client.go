package types

import dockertypes "github.com/docker/docker/api/types"

type RedisGetOpts struct {
	Key       string
	ServiceID string
	Node      Address
}

type RedisSetOpts struct {
	Key       string
	Value     string
	ServiceID string
	Node      Address
}

// StartOpts ...
type StartOpts struct {
	CStartOpts  dockertypes.ContainerStartOptions
	ContainerID string
}

type StopOpts struct {
	ContainerID string
	Timeout     string
}

type ListOpts struct {
	dockertypes.ContainerListOptions
}
type MigrateOpts struct {
	Address
	ServiceID     string
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}
