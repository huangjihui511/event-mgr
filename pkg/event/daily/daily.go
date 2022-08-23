package daily

import (
	"context"
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/interfaces"
	"sync"

	cronV3 "github.com/robfig/cron/v3"
)

var (
	_ eventInterface.Interface = Daily{}
)

type Daily struct {
	spec    string
	once    sync.Once
	c       chan interface{}
	watcher watcherInterface.Interface
}

// Chan implements interfaces.Interface
func (d Daily) Chan(ctx context.Context) <-chan interface{} {
	d.once.Do(func() {
		d.c = make(chan interface{})
		c := newWithSeconds()
		_, err := c.AddFunc(d.spec, func() {
			d.c <- struct{}{}
		})
		if err != nil {
			panic(err)
		}
		c.Start()
	})
	return d.c
}

// Watcher implements interfaces.Interface
func (d Daily) Watcher() watcherInterface.Interface {
	return d.watcher
}

func NewDaily(spec string, watcher watcherInterface.Interface) eventInterface.Interface {
	return &Daily{
		spec:    spec,
		watcher: watcher,
	}
}

func newWithSeconds() *cronV3.Cron {
	secondParser := cronV3.NewParser(cronV3.Second | cronV3.Minute |
		cronV3.Hour | cronV3.Dom | cronV3.Month | cronV3.DowOptional | cronV3.Descriptor)
	return cronV3.New(cronV3.WithParser(secondParser), cronV3.WithChain())
}
