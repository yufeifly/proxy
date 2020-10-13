package model

import (
	"github.com/yufeifly/proxy/config"
	"sync"
)

type Log struct {
	Last     bool
	LogQueue []string // log container
}

func newLog() *Log {
	return &Log{
		Last:     false,
		LogQueue: []string{},
	}
}

type Logger struct {
	Log
	//logQueue  []string // log container
	Count     int // current log entry count
	Capacity  int // size of a log page
	TotalSend int
	sync.Mutex
}

func NewLogger() *Logger {
	return &Logger{
		Log: Log{
			Last:     false,
			LogQueue: []string{},
		},
		Count:     0,
		Capacity:  config.Capacity,
		TotalSend: 0,
		Mutex:     sync.Mutex{},
	}
}

func (l *Logger) ClearQueue() {
	l.LogQueue = []string{}
}

func (l *Logger) SetLastFlag() {
	l.Last = true
}
