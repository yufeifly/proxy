package model

type Cluster struct {
	Master Node   `json:"proxy"`
	Worker []Node `json:"worker"`
}

type Node struct {
	Address
}
