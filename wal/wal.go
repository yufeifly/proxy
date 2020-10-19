package wal

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/scheduler"
)

var logger *model.Logger

func init() {
	logger = model.NewLogger()
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

// LockAndGetSentConsumed return sent and consumed
func LockAndGetSentConsumed() (int, int) {
	logger.Lock()
	sent := logger.TotalSend
	consumed := logger.TotalConsumed
	return sent, consumed
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

func ConsumedAdder() {
	logger.Lock()
	logger.TotalConsumed++
	logger.Unlock()
}
