package client

import (
	dockertypes "github.com/docker/docker/api/types"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/api/types/svc"
)

// APIClient ...
type APIClient interface {
	ContainerAPIClient
	ServiceAPIClient
	MigrationAPIClient
}

// ContainerAPIClient ...
type ContainerAPIClient interface {
	ContainerList(options types.ListOpts) ([]dockertypes.Container, error)
	ContainerStart(options types.StartOpts) error
	StopContainer(options types.StopOpts) error
}

// ServiceAPIClient ...
type ServiceAPIClient interface {
	AddService(service svc.ServiceOpts) error
}

// MigrationAPIClient ...
type MigrationAPIClient interface {
	SendMigrate(options types.MigrateOpts) error
}
