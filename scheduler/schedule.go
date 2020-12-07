package scheduler

import (
	"github.com/yufeifly/proxy/cusErr"
	"sync"
)

var DefaultScheduler *Scheduler

func init() {
	DefaultScheduler = NewScheduler()
}

type Scheduler struct {
	Map sync.Map
}

// NewScheduler new a scheduler
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// Default default scheduler
func Default() *Scheduler {
	return DefaultScheduler
}

// AddService add a service to scheduler
func (s *Scheduler) AddService(sID string, service *Service) {
	s.Map.Store(sID, service)
}

// GetService get service from scheduler
func (s *Scheduler) GetService(serviceID string) (*Service, error) {
	serviceP, ok := s.Map.Load(serviceID)
	if !ok {
		return nil, cusErr.ErrServiceNotFound
	}
	service, _ := serviceP.(*Service)
	return service, nil
}

// DeleteService ...
func (s *Scheduler) DeleteService(serviceID string) {
	s.Map.Delete(serviceID)
}

// ListService list all services of a scheduler
func (s *Scheduler) ListService() (services []*Service) {
	s.Map.Range(func(key, value interface{}) bool {
		ser, _ := value.(*Service)
		services = append(services, ser)
		return true
	})
	return
}
