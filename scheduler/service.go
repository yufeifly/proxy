package scheduler

import (
	"github.com/yufeifly/proxy/config"
	"github.com/yufeifly/proxy/model"
)

type Service struct {
	ID   string        // service id
	Node model.Address // the node that service exists
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
		ID: "service1",
		NodeAddr: model.Address{
			IP:   "127.0.0.1",
			Port: config.DefaultMigratorListeningPort,
		},
	}
	register(opts1)
	//DefaultScheduler.AddService(NewService(opts1))

	opts2 := model.ServiceOpts{
		ID: "service2",
		NodeAddr: model.Address{
			IP:   "127.0.0.1",
			Port: config.DefaultMigratorListeningPort,
		},
	}
	register(opts2)
	//DefaultScheduler.AddService(NewService(opts2))
}

func register(opts model.ServiceOpts) {
	DefaultScheduler.AddService(NewService(opts))
	// send to src node
}
