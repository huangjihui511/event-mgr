package event

import (
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	notifyInterface "huangjihui511/event-mgr/pkg/notify/interfaces"
	"huangjihui511/event-mgr/pkg/watcher/devops"
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
		NewTimer(time.Minute*30, scb.NewWatcherExchangeRatio("SCB Watcher")),
		NewTimer(time.Minute, scb.NewWatcherExchangeRatioLowerBuyRatio("SCB Low Bound Watcher", 6.78)),
		NewTimer(time.Second, devops.NewWatcherBorn()),
	}
}
