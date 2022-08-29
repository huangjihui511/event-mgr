package event

import (
	"huangjihui511/event-mgr/pkg/event/daily"
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	"huangjihui511/event-mgr/pkg/event/timer"
	"huangjihui511/event-mgr/pkg/watcher/devops"
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"huangjihui511/event-mgr/pkg/watcher/zsb"
	"time"
)

type RouterItem struct {
	EventName string
	Email     string
}

var (
	RouterItems = map[RouterItem]bool{
		{
			EventName: "SCB Watcher Daily",
			Email:     EmailJihui,
		}: true,
		{
			EventName: "ZSB Watcher Daily",
			Email:     EmailJihui,
		}: true,
	}
	EmailJihui   = "717655909@qq.com"
	targetEmails = []string{
		// EmailJihui,
	}
)

func registerEvents() {
	events = []eventInterface.Interface{
		timer.NewTimer(time.Second, devops.NewWatcherBorn()),
		timer.NewTimer(time.Minute*30, scb.NewWatcherExchangeRatio("SCB Watcher")),
		timer.NewTimer(time.Minute, scb.NewWatcherExchangeRatioLowerBuyRatio("SCB Low Bound Watcher 6.8", 6.8)),
		timer.NewTimer(time.Minute*30, zsb.NewWatcherExchangeRatio("ZSB Watcher")),
		timer.NewTimer(time.Minute, zsb.NewWatcherExchangeRatioLowerBuyRatio("ZSB Low Bound Watcher 6.8", 6.8)),
		daily.NewDaily("* 0 10 * * ?", scb.NewWatcherExchangeRatio("SCB Watcher Daily")),
		daily.NewDaily("* 0 10 * * ?", zsb.NewWatcherExchangeRatio("ZSB Watcher Daily")),
	}
}
