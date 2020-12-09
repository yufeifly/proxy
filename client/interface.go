package client

import (
	dockertypes "github.com/docker/docker/api/types"
	"github.com/yufeifly/proxy/api/logger"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/api/types/svc"
)

type APIClient interface {
	ContainerAPIClient
	ServiceAPIClient
	LogAPIClient
	RedisAPIClient
	MigrationAPIClient
}

type ContainerAPIClient interface {
	ContainerList(options types.ListOpts) ([]dockertypes.Container, error)
	ContainerStart(options types.StartOpts) error
	StopContainer(options types.StopOpts) error
}

type ServiceAPIClient interface {
	AddService(service svc.ServiceOpts) error
}

type LogAPIClient interface {
	SendLog(logWithID logger.LogWithServiceID) error
}

type RedisAPIClient interface {
	RedisGet(options types.RedisGetOpts) (string, error)
	RedisSet(options types.RedisSetOpts) error
}

type MigrationAPIClient interface {
	SendMigrate(options types.MigrateOpts) error
}
