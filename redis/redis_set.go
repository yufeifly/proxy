package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/cusErr"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/ticket"
	"github.com/yufeifly/proxy/wal"
)

// Set set kv pair to redis service
func Set(service string, key, val string) error {
	// if get ticket
	token := ticket.T.GetTicket()
	if token == ticket.ShutWrite {
		return cusErr.ErrServiceNotAvailable
	}
	if token == ticket.Logging {
		writeLog(key, val)
	}
	// get service
	ser, err := scheduler.DefaultScheduler.GetService(service)
	if err != nil {
		return err
	}
	// do set
	err = doSetKV(ser.ServiceCli, key, val)
	if err != nil {
		return err
	}
	return nil
}

//
func doSetKV(cli *redis.Client, key, val string) error {
	err := cli.Set(context.Background(), key, val, 0).Err()
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": val,
	}).Info("pair set")
	return nil
}

//
func writeLog(key, val string) error {
	logrus.Warn("operation send to message queue")
	str := []string{key, val}
	strJson, _ := json.Marshal(str)
	err := wal.DataLog(string(strJson))
	return err
}
