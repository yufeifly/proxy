package ticket

import "sync"

var defaultTicket *ticket

const (
	normal    = 0
	Logging   = 1
	ShutWrite = 2
)

// Ticket ...
type Ticket interface {
	Get() int
	Set(v int)
	UnSet()
}

type ticket struct {
	token int // 0 means , 1 means start logging, 2 means shut write
	rw    sync.RWMutex
}

// InitTicket ...
func InitTicket() {
	defaultTicket = NewTicket()
}

// NewTicket ...
func NewTicket() *ticket {
	return &ticket{
		token: 0,
		rw:    sync.RWMutex{},
	}
}

// Default default ticket
func Default() Ticket {
	return defaultTicket
}

// Get get value of ticket
func (t *ticket) Get() int {
	var ret int
	t.rw.RLock()
	ret = t.token
	t.rw.RUnlock()
	return ret
}

// Set set value of ticket
func (t *ticket) Set(v int) {
	t.rw.Lock()
	t.token = v
	t.rw.Unlock()
}

// UnSet restore to normal mode
func (t *ticket) UnSet() {
	t.rw.Lock()
	t.token = normal
	t.rw.Unlock()
}
