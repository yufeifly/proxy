package scheduler

import (
	"github.com/yufeifly/proxy/cuserr"
	"sync"
)

var defaultScheduler *scheduler

// Scheduler ...
type Scheduler interface {
	AddService(id string, service *Service) // id means proxyServiceID
	GetService(id string) (*Service, error)
	ListService() (services []*Service)
}

type scheduler struct {
	Map sync.Map
}

// InitScheduler init scheduler
func InitScheduler() {
	defaultScheduler = NewScheduler()
}

// NewScheduler new a scheduler
func NewScheduler() *scheduler {
	return &scheduler{}
}

// Default get default scheduler interface
func Default() Scheduler {
	return defaultScheduler
}

// AddService add a service to scheduler
func (s *scheduler) AddService(id string, service *Service) {
	s.Map.Store(id, service)
}

// GetService get service from scheduler
func (s *scheduler) GetService(id string) (*Service, error) {
	serviceP, ok := s.Map.Load(id)
	if !ok {
		return nil, cuserr.ErrServiceNotFound
	}
	service, _ := serviceP.(*Service)
	return service, nil
}

// DeleteService ...
func (s *scheduler) DeleteService(id string) {
	s.Map.Delete(id)
}

// ListService list all services of a scheduler
func (s *scheduler) ListService() (services []*Service) {
	s.Map.Range(func(key, value interface{}) bool {
		ser, _ := value.(*Service)
		services = append(services, ser)
		return true
	})
	return
}
