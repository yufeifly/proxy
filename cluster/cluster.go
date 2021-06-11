package cluster

import "github.com/yufeifly/proxy/api/types"

// Cluster ...
type Cluster interface {
	// get proxy of the cluster
	GetProxy() Node
	// get workers of the cluster
	GetWorkers() []Node
}

type cluster struct {
	Master  Node   `json:"proxy"`
	Workers []Node `json:"worker"`
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
func (c *cluster) GetProxy() Node {
	return c.Master
}

// Workers ...
func (c *cluster) GetWorkers() []Node {
	return c.Workers
}
