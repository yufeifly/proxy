package model

type Cluster struct {
	Master Node   `json:"proxy"`
	Worker []Node `json:"worker"`
}

type Node struct {
	Address
}

// Cluster.GetProxy ...
func (c *Cluster) GetProxy() Node {
	return c.Master
}

// Cluster.GetWorkers ...
func (c *Cluster) GetWorkers() []Node {
	return c.Worker
}
