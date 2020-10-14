package wal

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
)

var logger *model.Logger
var consumer *Consumed

func init() {
	logger = model.NewLogger()
	consumer = NewConsumed()
}

// DataLog
func DataLog(data string) error {
	logger.Lock()
	logger.Count++
	logger.LogQueue = append(logger.LogQueue, data)
	if logger.Count == logger.Capacity {
		fmt.Println("send to dst by goroutine")
		cli := client.Client{}
		cli.SendLog(logger.Log)
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
func SendLastLog() error {
	logrus.Info("send the last log")
	cli := client.Client{}

	logger.Lock()
	defer logger.Unlock()

	logger.SetLastFlag()
	err := cli.SendLog(logger.Log)
	if err != nil {
		return err
	}
	//logger.TotalSend++
	//logger.Count = 0
	//if len(logger.LogQueue) > 0 {
	//	logger.TotalSend++
	//	logger.Count--
	//}

	logrus.Info("SetLastLog finished")
	// ? set an empty log, to let dst consumer goroutine to stop
	return nil
}
