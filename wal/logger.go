package wal

import "sync"

type Logger struct {
	logQueue  []string // log container
	count     int      // current log entry count
	capacity  int      // size of a log page
	TotalSend int
	sync.Mutex
}

func NewLogger() *Logger {
	return &Logger{
		logQueue:  []string{},
		count:     0,
		capacity:  Capacity,
		TotalSend: 0,
		Mutex:     sync.Mutex{},
	}
}

func (l *Logger) ClearQueue() {
	l.logQueue = []string{}
}
