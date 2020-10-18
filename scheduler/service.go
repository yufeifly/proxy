package scheduler

import (
	"github.com/yufeifly/proxy/config"
	"github.com/yufeifly/proxy/model"
)

type Service struct {
	ID     string        // service id
	Node   model.Address // the node that service exists
	Shadow model.Address // node that may replace the origin node, useful in migration
}

func init() {
	PseudoRegister()
}

// NewService new a storage service, keep it in map
func NewService(opts model.ServiceOpts) *Service {
	return &Service{
		ID:   opts.ID,
		Node: opts.NodeAddr,
	}
}

// PseudoRegister register services
func PseudoRegister() {
	opts1 := model.ServiceOpts{
		ID: "service.A1",
		NodeAddr: model.Address{
			IP:   "127.0.0.1",
			Port: config.DefaultMigratorListeningPort,
		},
	}
	register("service1", opts1)

	opts2 := model.ServiceOpts{
		ID: "service.B1",
		NodeAddr: model.Address{
			IP:   "127.0.0.1",
			Port: config.DefaultMigratorListeningPort,
		},
	}
	register("service2", opts2)
}

func register(service string, opts model.ServiceOpts) {
	DefaultScheduler.AddService(service, NewService(opts))
}

func (s *Service) AddShadow(addr model.Address) {
	s.Shadow = addr
}
