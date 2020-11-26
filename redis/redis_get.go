package redis

import (
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/scheduler"
)

// Get redis get value by key
func Get(ProxyService string, key string) (string, error) {
	// get service
	service, err := scheduler.Default().GetService(ProxyService)
	if err != nil {
		return "", err
	}
	// send get request
	getOpts := model.RedisGetOpts{
		Key:       key,
		ServiceID: service.ID,
		Node:      service.Node,
	}
	cli := client.Client{}
	val, err := cli.RedisGet(getOpts)
	if err != nil {
		return "", err
	}
	return val, err
}
