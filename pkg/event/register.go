package event

import (
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	notifyInterface "huangjihui511/event-mgr/pkg/notify/interfaces"
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"time"
)

var (
	events       []eventInterface.Interface
	notifyEmail  notifyInterface.EmailInterface
	targetEmails = []string{
		"717655909@qq.com",
	}
)

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
