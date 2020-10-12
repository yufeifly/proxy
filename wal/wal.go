package wal

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
)

var logger *Logger
var consumer *Consumed

func init() {
	logger = NewLogger()
	consumer = NewConsumed()
}

// DataLog
func DataLog(data string) error {
	logger.Lock()
	logger.count++
	logger.logQueue = append(logger.logQueue, data)
	if logger.count == logger.capacity {
		fmt.Println("send to dst by goroutine")
		cli := client.Client{}
		cli.SendLog(logger.logQueue)
		logger.TotalSend++
		logger.ClearQueue()
		logger.count = 0
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

func SendLogEntry() {
	logrus.Info("send leftover logs")
	cli := client.Client{}
	logger.Lock()
	cli.SendLog(logger.logQueue)
	logger.TotalSend++
	logger.count = 0
	logger.Unlock()
}
