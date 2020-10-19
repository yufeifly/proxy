package wal

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/scheduler"
)

var logger *model.Logger
var consumer *Consumed

func init() {
	logger = model.NewLogger()
	consumer = NewConsumed()
}

// DataLog
func DataLog(ser *scheduler.Service, data string) error {
	logger.Lock()
	logger.Count++
	logger.LogQueue = append(logger.LogQueue, data)
	if logger.Count == logger.Capacity {
		// todo send to dst by goroutine
		cli := client.Client{
			Dest: ser.Shadow,
		}
		logWithID := model.LogWithServiceID{
			Log:            logger.Log,
			ProxyServiceID: ser.ProxyServiceID,
		}
		cli.SendLog(logWithID)
		logger.TotalSend++
		logger.ClearQueue()
		logger.Count = 0
	}
	logger.Unlock()
	return nil
}

func LockAndGetTotalSend() int {
	var ret int
	logger.Lock()
	ret = logger.TotalSend
	return ret
}

func UnlockLogger() {
	logger.Unlock()
}

// SendLastLog send the last log to dst
func SendLastLog(ProxyServiceID string, addr model.Address) error {
	logrus.Info("send the last log")
	cli := client.Client{
		Dest: addr,
	}

	logger.Lock()
	defer logger.Unlock()

	logger.SetLastFlag()
	logWithID := model.LogWithServiceID{
		Log:            logger.Log,
		ProxyServiceID: ProxyServiceID,
	}
	err := cli.SendLog(logWithID)
	if err != nil {
		return err
	}

	logrus.Infof("SetLastLog finished, ProxyServiceID: %v", ProxyServiceID)
	return nil
}
