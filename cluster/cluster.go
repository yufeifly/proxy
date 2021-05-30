package cluster

import "github.com/yufeifly/proxy/api/types"

// Cluster ...
type Cluster interface {
	// get proxy of the cluster
	Proxy() Node
	// get workers of the cluster
	Workers() []Node
}

type cluster struct {
	master  Node   `json:"proxy"`
	workers []Node `json:"worker"`
}

var defaultCluster cluster

// Node ...
type Node struct {
	types.Address
}

// Default return default cluster
func Default() Cluster {
	return &defaultCluster
}

// Proxy ...
func (c *cluster) Proxy() Node {
	return c.master
}

// Workers ...
func (c *cluster) Workers() []Node {
	return c.workers
}
