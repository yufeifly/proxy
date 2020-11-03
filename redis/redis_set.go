package redis

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/cusErr"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/ticket"
	"github.com/yufeifly/proxy/wal"
)

// Set set kv pair to redis service
func Set(ProxyService string, key, val string) error {
	// if get ticket
	token := ticket.Default().Get()
	if token == ticket.ShutWrite {
		return cusErr.ErrServiceNotAvailable
	}
	// get service
	service, err := scheduler.Default().GetService(ProxyService)
	if err != nil {
		return err
	}
	if token == ticket.Logging {
		writeLog(service, key, val)
	}
	// send set request
	cli := client.Client{}
	err = cli.RedisSet(service, key, val)
	if err != nil {
		return err
	}
	return nil
}

//
func writeLog(ser *scheduler.Service, key, val string) error {
	logrus.Warn("operation send to message queue")
	str := []string{key, val}
	strJson, _ := json.Marshal(str)
	err := wal.DataLog(ser, string(strJson))
	return err
}
