package ticket

import "sync"

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

// NewTicket return ticket interface
func NewTicket() Ticket {
	return &ticket{
		token: normal,
		rw:    sync.RWMutex{},
	}
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
