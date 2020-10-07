package redis

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/dal"
	"github.com/yufeifly/proxy/ticket"
)

func Set(key, val string) error {
	// if get ticket
	if !ticket.T.GetTicket() {
		err := dal.SetKV(key, val)
		if err != nil {
			return err
		}
	} else { // if not get ticket
		// send the op to message queue, (goroutine)
		logrus.Warn("op send to message queue")
		// do set
		err := dal.SetKV(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}
