package types

import dockertypes "github.com/docker/docker/api/types"

// RedisGetOpts ...
type RedisGetOpts struct {
	Key       string
	ServiceID string
	Node      Address
}

// RedisSetOpts ...
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

// StopOpts ...
type StopOpts struct {
	ContainerID string
	Timeout     string
}

// ListOpts ...
type ListOpts struct {
	dockertypes.ContainerListOptions
}

// MigrateOpts ...
type MigrateOpts struct {
	Address
	ServiceID     string
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}
