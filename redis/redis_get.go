package redis

import (
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/scheduler"
)

// Get redis get value by key
func Get(ProxyService string, key string) (string, error) {
	// get service
	service, err := scheduler.DefaultScheduler.GetService(ProxyService)
	if err != nil {
		return "", err
	}
	// send get request
	cli := client.Client{}
	val, err := cli.RedisGet(service, key)
	if err != nil {
		return "", err
	}
	return val, err
}
