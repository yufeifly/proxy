package logger

import (
	"sync"
)

const capacity = 10

// Log ...
type Log struct {
	Last     bool
	LogQueue []string // log container
}

// NewLog new a log
func NewLog() *Log {
	return &Log{
		Last:     false,
		LogQueue: []string{},
	}
}

// LogWithPSID ...
type LogWithPSID struct {
	ProxyServiceID string
	*Log
}

// NewLogWithServiceID new a log with service ID
func NewLogWithServiceID(proxyService string) *LogWithPSID {
	return &LogWithPSID{
		ProxyServiceID: proxyService,
		Log:            NewLog(),
	}
}

// Logger ...
type Logger struct {
	Count    int // current log entry count
	Capacity int // size of a log page
	Sent     int
	Consumed int
	sync.Mutex
	*LogWithPSID
}

// NewLogger ...
func NewLogger(proxyService string) *Logger {
	return &Logger{
		Count:       0,
		Capacity:    capacity,
		Sent:        0,
		Consumed:    0,
		Mutex:       sync.Mutex{},
		LogWithPSID: NewLogWithServiceID(proxyService),
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
