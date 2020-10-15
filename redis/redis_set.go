package redis

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/cusErr"
	"github.com/yufeifly/proxy/dal"
	"github.com/yufeifly/proxy/ticket"
	"github.com/yufeifly/proxy/wal"
)

// Set set kv pair to redis service
func Set(key, val string) error {
	// if get ticket
	token := ticket.T.GetTicket()
	if token == ticket.ShutWrite {
		return cusErr.ErrServiceNotAvailable
	}
	if token == ticket.Logging {
		writeLog(key, val)
	}
	// do redis set
	err := dal.SetKV(key, val)
	if err != nil {
		return err
	}
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
