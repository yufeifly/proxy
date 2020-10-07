package client

import (
	"context"
	"github.com/yufeifly/proxy/model"
)

type APIClient interface {
	CommonAPIClient
}

type CommonAPIClient interface {
	ContainerAPIClient
	ImageAPIClient
	MigrateAPIClient
}

type ContainerAPIClient interface {
}

type ImageAPIClient interface {
}

type MigrateAPIClient interface {
	Migrate(ctx context.Context, options model.MigrateOpts) error
}
