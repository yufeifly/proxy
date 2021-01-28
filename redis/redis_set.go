package redis

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/cuserr"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/ticket"
)

// Set set kv pair to redis service
func Set(proxyService string, key, val string) error {
	// get service
	service, err := scheduler.Default().GetService(proxyService)
	if err != nil {
		logrus.Errorf("GetService failed, err: %v", err)
		return err
	}
	token := service.Ticket().Get()
	// if ticket = ShutWrite
	if token == ticket.ShutWrite {
		return cuserr.ErrServiceNotAvailable
	}
	// if ticket = Logging, log it
	if token == ticket.Logging {
		err := logRecord(service, key, val)
		if err != nil {
			logrus.Errorf("redis.Set logRecord failed, err: %v", err)
			return err
		}
	}
	// send set request
	opts := types.RedisSetOpts{
		Key:       key,
		Value:     val,
		ServiceID: service.ID,
		Node:      service.Node,
	}
	cli := client.DefaultClient()
	err = cli.RedisSet(opts)
	if err != nil {
		return err
	}
	return nil
}

// logRecord logs a record
func logRecord(service *scheduler.Service, key, val string) error {
	logrus.Warn("redis.logRecord logging operation")
	kv := []string{key, val}
	kvJSON, err := json.Marshal(kv)
	if err != nil {
		return err
	}
	err = service.LogDataInJSON(string(kvJSON))
	return err
}
