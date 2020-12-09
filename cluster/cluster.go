package cluster

import "github.com/yufeifly/proxy/api/types"

// Cluster ...
type Cluster struct {
	Master Node   `json:"proxy"`
	Worker []Node `json:"worker"`
}

// Node ...
type Node struct {
	types.Address
}

// GetProxy ...
func (c *Cluster) GetProxy() Node {
	return c.Master
}

// GetWorkers ...
func (c *Cluster) GetWorkers() []Node {
	return c.Worker
}
