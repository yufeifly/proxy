package scheduler

import (
	"github.com/go-redis/redis/v8"
	"github.com/yufeifly/proxy/model"
)

type Service struct {
	ID          string        // service id
	Node        model.Address // the node that service exists
	ContainerID string        // the worker container
	ServiceCli  *redis.Client // redis connection
}

func init() {
	PseudoRegister()
}

// NewService new a storage service, keep it in map
func NewService(opts model.ServiceOpts) *Service {
	return &Service{
		ID:          opts.ID,
		Node:        opts.NodeAddr,
		ContainerID: opts.Container,
		ServiceCli: redis.NewClient(&redis.Options{
			Addr:     opts.NodeAddr.IP + ":" + opts.NodeAddr.Port,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

// PseudoRegister register services
func PseudoRegister() {
	opts1 := model.ServiceOpts{
		ID: "service1",
		NodeAddr: model.Address{
			IP:   "127.0.0.1",
			Port: "6380",
		},
		Container: "19571cac86b3",
	}
	DefaultScheduler.AddService(NewService(opts1))

	opts2 := model.ServiceOpts{
		ID: "service2",
		NodeAddr: model.Address{
			IP:   "127.0.0.1",
			Port: "6666",
		},
		Container: "30860d58aebb",
	}
	DefaultScheduler.AddService(NewService(opts2))
}
