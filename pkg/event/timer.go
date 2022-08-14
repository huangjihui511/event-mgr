package event

import (
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/interfaces"
	"sync"
	"time"
)

var (
	_ eventInterface.Interface = Timer{}
)

type Timer struct {
	duration         time.Duration
	watcherInterface watcherInterface.Interface
	once             sync.Once
	c                chan interface{}
}

func NewTimer(d time.Duration, w watcherInterface.Interface) eventInterface.Interface {
	return Timer{
		duration:         d,
		watcherInterface: w,
	}
}

func (t Timer) Chan() <-chan interface{} {
	t.once.Do(func() {
		t.c = make(chan interface{}, 0)
		go func() {
			for tt := range time.After(t.duration) {
				t.c <- tt
			}
		}()
	})
	return t.c
}

func (t Timer) Watcher() watcherInterface.Interface {
	return t.watcherInterface
}
