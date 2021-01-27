package redis

import (
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/scheduler"
)

// Delete delete
func Delete(ProxyService string, key string) error {
	// get service
	service, err := scheduler.Default().GetService(ProxyService)
	if err != nil {
		return err
	}
	// send get request
	getOpts := types.RedisDeleteOpts{
		Key:       key,
		ServiceID: service.ID,
		Node:      service.Node,
	}
	cli := client.NewClient(getOpts.Node)
	err = cli.RedisDelete(getOpts)
	if err != nil {
		return err
	}
	return err
}
