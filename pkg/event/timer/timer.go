package timer

import (
	"context"
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

func (t Timer) Chan(ctx context.Context) <-chan interface{} {
	t.once.Do(func() {
		t.c = make(chan interface{})
		go func() {
			for {
				var tc time.Time
				select {
				case tc = <-time.After(t.duration):
					t.c <- tc
				case <-ctx.Done():
					return
				}
			}
		}()
	})
	return t.c
}

func (t Timer) Watcher() watcherInterface.Interface {
	return t.watcherInterface
}
