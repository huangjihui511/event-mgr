package event

import (
	"context"
	eventInterface "huangjihui511/event-mgr/pkg/event/event_interface"
	"huangjihui511/event-mgr/pkg/logs"
	"huangjihui511/event-mgr/pkg/notify"
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"time"
)

var (
	events []eventInterface.Interface
)

func StartMgr(ctx context.Context) {
	registerEvents()
	startEvents(ctx)
}

func registerEvents() {
	events = []eventInterface.Interface{
		Timer{
			watcherInterface: scb.ExchangeRatioWatcher{
				Name_: "scb watcher",
			},
			duration: time.Second * 5,
		},
	}
}

func startEvents(ctx context.Context) {
	for _, e := range events {
		go func(ev eventInterface.Interface) {
			c := ev.Chan()
			for {
				select {
				case <-c:
					logs.Logger.Infof("Call watcher %s", ev.Watcher().Name())
					r := ev.Watcher().Call(ctx)
					if !r.IsNotify() {
						continue
					}
					notify.SendToEmail(r.Subject(), r.Info())
				case <-ctx.Done():
					logs.Logger.Infof("Stop watcher %s", ev.Watcher().Name())
					return
				}
			}
		}(e)
	}
}
