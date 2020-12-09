package cluster

import "github.com/yufeifly/proxy/api/types"

var defaultCluster cluster

// Cluster ...
type Cluster interface {
	GetProxy() Node
	GetWorkers() []Node
}

type cluster struct {
	Master  Node   `json:"proxy"`
	Workers []Node `json:"worker"`
}

// Node ...
type Node struct {
	types.Address
}

// DefaultCluster return default cluster
func DefaultCluster() Cluster {
	return &defaultCluster
}

// GetProxy ...
func (c *cluster) GetProxy() Node {
	return c.Master
}

// GetWorkers ...
func (c *cluster) GetWorkers() []Node {
	return c.Workers
}
