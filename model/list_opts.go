package model

import "github.com/docker/docker/api/types"

type ListReqOpts struct {
	Address
	types.ContainerListOptions
}
