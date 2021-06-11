package scheduler

import (
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/api/types/svc"
	"github.com/yufeifly/proxy/cluster"
	"github.com/yufeifly/proxy/config"
)

// Service ...
type Service struct {
	ID             string // service id
	ProxyServiceID string
	Node           types.Address // the node that service exists
	MigTarget      types.Address // node that may replace the origin node, useful in migration
}

// NewService new a storage service, keep it in map
func NewService(opts svc.ServiceOpts) *Service {
	return &Service{
		ID:             opts.ID,
		ProxyServiceID: opts.ProxyServiceID,
		Node:           opts.NodeAddr,
		MigTarget:      types.Address{},
	}
}

// PseudoRegister register services
func PseudoRegister() {
	proxyIP := cluster.Default().GetProxy().IP
	opts1 := svc.ServiceOpts{
		ID:             "service1.1",
		ProxyServiceID: "service1",
		NodeAddr: types.Address{
			IP:   proxyIP,
			Port: config.DefaultMigratorListeningPort,
		},
	}
	DefaultRegister("service1", opts1)

	opts2 := svc.ServiceOpts{
		ID:             "service2.1",
		ProxyServiceID: "service2",
		NodeAddr: types.Address{
			IP:   proxyIP,
			Port: config.DefaultMigratorListeningPort,
		},
	}
	DefaultRegister("service2", opts2)
}

// DefaultRegister register a service to defaultScheduler
func DefaultRegister(ProxyService string, opts svc.ServiceOpts) {
	Default().AddService(ProxyService, NewService(opts))
}

// AddMigTarget ...
func (s *Service) AddMigTarget(addr types.Address) {
	s.MigTarget.IP = addr.IP
	s.MigTarget.Port = addr.Port
}
