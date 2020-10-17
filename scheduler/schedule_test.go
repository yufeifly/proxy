package scheduler

import (
	"fmt"
	"github.com/yufeifly/proxy/model"
	"testing"
)

func TestScheduler_AddService(t *testing.T) {
	service := &Service{
		ID: "service1",
		Node: model.Address{
			IP:   "127.0.0.1",
			Port: "6789",
		},
		ContainerID: "123456",
	}
	DefaultScheduler.AddService(service)
	get, _ := DefaultScheduler.GetService(service.ID)
	fmt.Printf("get serive: %v\n", get)
}
