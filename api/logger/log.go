package logger

import (
	"github.com/yufeifly/proxy/config"
	"sync"
)

type Log struct {
	Last     bool
	LogQueue []string // log container
}

func NewLog() *Log {
	return &Log{
		Last:     false,
		LogQueue: []string{},
	}
}

type LogWithServiceID struct {
	*Log
	ProxyServiceID string
}

func NewLogWithServiceID(proxyService string) *LogWithServiceID {
	return &LogWithServiceID{
		Log:            NewLog(),
		ProxyServiceID: proxyService,
	}
}

type Logger struct {
	*LogWithServiceID
	Count         int // current log entry count
	Capacity      int // size of a log page
	TotalSend     int
	TotalConsumed int
	sync.Mutex
}

func NewLogger(proxyService string) *Logger {
	return &Logger{
		LogWithServiceID: NewLogWithServiceID(proxyService),
		Count:            0,
		Capacity:         config.Capacity,
		TotalSend:        0,
		TotalConsumed:    0,
		Mutex:            sync.Mutex{},
	}
}

// ClearQueue clear data queue of the queue
func (l *Logger) ClearQueue() {
	l.LogQueue = l.LogQueue[:0]
}

// SetLastFlag set last flag of the log, which means logging will end
func (l *Logger) SetLastFlag() {
	l.Last = true
}
