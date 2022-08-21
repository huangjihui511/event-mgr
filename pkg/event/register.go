package event

import (
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	"huangjihui511/event-mgr/pkg/notify"
	notifyInterface "huangjihui511/event-mgr/pkg/notify/interfaces"
	"huangjihui511/event-mgr/pkg/watcher/devops"
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"huangjihui511/event-mgr/pkg/watcher/zsb"
	"time"
)

var (
	events       []eventInterface.Interface
	notifyEmail  notifyInterface.EmailInterface
	targetEmails = []string{
		"717655909@qq.com",
	}
)

func init() {
	notifyEmail = notify.EmailSender{
		EmailMeta: notifyInterface.EmailMetaQQ,
	}
}

func registerEvents() {
	events = []eventInterface.Interface{
		NewTimer(time.Second, devops.NewWatcherBorn()),
		NewTimer(time.Minute*30, scb.NewWatcherExchangeRatio("SCB Watcher")),
		NewTimer(time.Minute, scb.NewWatcherExchangeRatioLowerBuyRatio("SCB Low Bound Watcher 6.8", 6.8)),
		NewTimer(time.Minute*30, zsb.NewWatcherExchangeRatio("ZSB Watcher")),
		NewTimer(time.Minute, zsb.NewWatcherExchangeRatioLowerBuyRatio("ZSB Low Bound Watcher 6.8", 6.8)),
	}
}
