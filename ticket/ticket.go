package ticket

import "sync"

var T *Ticket

const (
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

func (t *Ticket) GetTicket() int {
	var ret int
	t.rw.RLock()
	ret = t.token
	t.rw.RUnlock()
	return ret
}

func (t *Ticket) SetTicket(v int) {
	t.rw.Lock()
	t.token = v
	t.rw.Unlock()
}

func (t *Ticket) UnSet() {
	t.rw.Lock()
	t.token = 0
	t.rw.Unlock()
}
