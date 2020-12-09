package cluster

import "github.com/yufeifly/proxy/api/types"

type Cluster struct {
	Master Node   `json:"proxy"`
	Worker []Node `json:"worker"`
}

type Node struct {
	types.Address
}

// Cluster.GetProxy ...
func (c *Cluster) GetProxy() Node {
	return c.Master
}

// Cluster.GetWorkers ...
func (c *Cluster) GetWorkers() []Node {
	return c.Worker
}
