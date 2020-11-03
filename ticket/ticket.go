package ticket

import "sync"

var T *Ticket

const (
	Normal    = 0
	Logging   = 1
	ShutWrite = 2
)

func init() {
	T = NewTicket()
}

type Ticket struct {
	token int // 0 means , 1 means start logging, 2 means shut write
	rw    sync.RWMutex
}

func NewTicket() *Ticket {
	return &Ticket{
		token: 0,
		rw:    sync.RWMutex{},
	}
}

func Default() *Ticket {
	return T
}

// Get get value of ticket
func (t *Ticket) Get() int {
	var ret int
	t.rw.RLock()
	ret = t.token
	t.rw.RUnlock()
	return ret
}

// Set set value of ticket
func (t *Ticket) Set(v int) {
	t.rw.Lock()
	t.token = v
	t.rw.Unlock()
}

// UnSet restore to normal mode
func (t *Ticket) UnSet() {
	t.rw.Lock()
	t.token = Normal
	t.rw.Unlock()
}
