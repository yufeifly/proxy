package redis

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/cusErr"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/ticket"
)

// Set set kv pair to redis service
func Set(ProxyService string, key, val string) error {

	token := ticket.Default().Get()
	// if ticket == ShutWrite
	if token == ticket.ShutWrite {
		return cusErr.ErrServiceNotAvailable
	}
	// get service
	service, err := scheduler.Default().GetService(ProxyService)
	if err != nil {
		logrus.Errorf("GetService failed, err: %v", err)
		return err
	}
	// if ticket == Logging, log it
	if token == ticket.Logging {
		logRecord(service, key, val)
	}
	// send set request
	opts := model.RedisSetOpts{
		Key:       key,
		Value:     val,
		ServiceID: service.ID,
		Node:      service.Node,
	}
	cli := client.Client{}
	err = cli.RedisSet(opts)
	if err != nil {
		return err
	}
	return nil
}

// logRecord logs a record
func logRecord(service *scheduler.Service, key, val string) error {
	logrus.Warn("logging operation")
	kv := []string{key, val}
	kvJson, err := json.Marshal(kv)
	if err != nil {
		return err
	}
	err = service.LogDataInJson(string(kvJson))
	return err
}
