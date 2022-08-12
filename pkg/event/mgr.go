package event

import (
	"context"
	eventInterface "huangjihui511/event-mgr/pkg/event/event_interface"
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"time"
)

var (
	events []eventInterface.Interface
)

func StartMgr() {
	registerEvents()
	startEvents()
}

func registerEvents() {
	events = []eventInterface.Interface{
		Timer{
			watcherInterface: scb.ExchangeRatioWatcher{},
			duration:         time.Second * 5,
		},
	}
}

func startEvents() {
	for _, e := range events {
		go func(ev eventInterface.Interface) {
			c := ev.Chan()
			for range c {
				r := ev.Watcher().Call(context.TODO())
				if !r.IsNotify() {
					continue
				}
			}
		}(e)
	}
}
