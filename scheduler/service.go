package scheduler

import (
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/api/types/svc"
	"github.com/yufeifly/proxy/cluster"
	"github.com/yufeifly/proxy/config"
	"sync"
)

// Service ...
type Service struct {
	CName    string
	SID      string
	NodeLock sync.RWMutex
	Node     types.Address // the node that service exists
}

// NewService new a storage service, keep it in map
func NewService(opts svc.ServiceOpts) *Service {
	return &Service{
		CName: opts.CName,
		SID:   opts.SID,
		Node:  opts.NodeAddr,
	}
}

// UpdateNode ...
func (s *Service) UpdateNode(newNode types.Address) {
	s.NodeLock.Lock()
	s.Node = newNode
	s.NodeLock.Unlock()
}

func (s *Service) GetNode() types.Address {
	s.NodeLock.RLock()
	defer s.NodeLock.RUnlock()
	return s.Node
}

// PseudoRegister register services
func PseudoRegister() {
	proxyIP := cluster.Default().GetProxy().IP
	opts := []svc.ServiceOpts{
		{
			CName: "s1.c1",
			SID:   "s1",
			NodeAddr: types.Address{
				IP:   proxyIP,
				Port: config.DefaultMigratorListeningPort,
			},
		},
		{
			CName: "s1.c1",
			SID:   "s2",
			NodeAddr: types.Address{
				IP:   proxyIP,
				Port: config.DefaultMigratorListeningPort,
			},
		},
	}

	for _, opt := range opts {
		DefaultRegister(opt.CName, opt)
	}
}

// DefaultRegister register a service to defaultScheduler
func DefaultRegister(cid string, opts svc.ServiceOpts) {
	Default().AddService(cid, NewService(opts))
}
