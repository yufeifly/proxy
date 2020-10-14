package wal

import "sync"

type Consumed struct {
	Total int
	sync.Mutex
}

func NewConsumed() *Consumed {
	return &Consumed{
		Total: 0,
		Mutex: sync.Mutex{},
	}
}

func ConsumedAddOne() {
	consumer.Lock()
	consumer.Total++
	consumer.Unlock()
}

func LockAndGetTotalConsumed() int {
	var ret int
	consumer.Lock()
	ret = consumer.Total
	return ret
}

func UnlockConsumer() {
	consumer.Unlock()
}
