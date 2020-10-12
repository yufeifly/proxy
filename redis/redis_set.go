package redis

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/cusErr"
	"github.com/yufeifly/proxy/dal"
	"github.com/yufeifly/proxy/ticket"
)

// Set set kv pair to redis service
func Set(key, val string) error {
	// if get ticket
	token := ticket.T.GetTicket()
	if token == ticket.ShutWrite {
		return cusErr.ErrServiceNotAvailable
	}
	if token == ticket.Logging {
		// send the op to message queue, (goroutine)
		logrus.Warn("operation send to message queue")
		// do redis set
		err := dal.SetKV(key, val)
		if err != nil {
			return err
		}
	}
	// do redis set
	err := dal.SetKV(key, val)
	if err != nil {
		return err
	}
	return nil
}
