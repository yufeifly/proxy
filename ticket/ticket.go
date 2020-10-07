package ticket

import "sync"

var T *Ticket

func init() {
	T = NewTicket()
}

type Ticket struct {
	token bool
	rw    sync.RWMutex
}

func NewTicket() *Ticket {
	return &Ticket{
		token: false,
		rw:    sync.RWMutex{},
	}
}

func (t *Ticket) GetTicket() bool {
	var ret bool
	t.rw.RLock()
	ret = t.token
	t.rw.RUnlock()
	return ret
}

func (t *Ticket) Set() {
	t.rw.Lock()
	t.token = true
	t.rw.Unlock()
}

func (t *Ticket) UnSet() {
	t.rw.Lock()
	t.token = false
	t.rw.Unlock()
}
