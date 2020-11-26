package model

type Cluster struct {
	Master Node   `json:"proxy"`
	Worker []Node `json:"worker"`
}

type Node struct {
	Address
}

func (c *Cluster) GetProxy() Node {
	return c.Master
}

func (c *Cluster) GetWorkers() []Node {
	return c.Worker
}
