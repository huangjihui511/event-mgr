package event

import (
	"context"
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	"huangjihui511/event-mgr/pkg/logs"
	"huangjihui511/event-mgr/pkg/notify"
	notifyInterface "huangjihui511/event-mgr/pkg/notify/interfaces"
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"time"
)

var (
	events      []eventInterface.Interface
	notifyEmail = notify.EmailSender{
		notifyInterface.EmailMetaQQ,
	}
	targetEmails = []string{
		"717655909@qq.com",
	}
)

func StartMgr(ctx context.Context) {
	registerEvents()
	startEvents(ctx)
}

func registerEvents() {
	events = []eventInterface.Interface{
		Timer{
			watcherInterface: scb.WatcherExchangeRatio{
				Name_: "scb watcher",
			},
			duration: time.Minute * 30,
		},
		Timer{
			watcherInterface: scb.WatcherExchangeRatioLowerBuyRatio{
				WatcherExchangeRatio: scb.WatcherExchangeRatio{
					Name_: "scb watcher",
				},
				LowBoundRatio: 6.78,
			},
			duration: time.Minute,
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
					for _, t := range targetEmails {
						err := notifyEmail.Send(t, r.Subject(), r.Msg())
						if err != nil {
							logs.Logger.Errorf("send email to %v failed: %s", t, err)
						}
					}
				case <-ctx.Done():
					logs.Logger.Infof("Stop watcher %s", ev.Watcher().Name())
					return
				}
			}
		}(e)
	}
}
