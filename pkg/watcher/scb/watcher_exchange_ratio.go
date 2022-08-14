package scb

import (
	"context"
	"fmt"
	"huangjihui511/event-mgr/pkg/utils"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/interfaces"
)

var (
	_ watcherInterface.Interface = WatcherExchangeRatio{}
	_ watcherInterface.Interface = WatcherExchangeRatioLowerBuyRatio{}
)

type WatcherExchangeRatio struct {
	watcherInterface.WatcherBase
}

func NewWatcherExchangeRatio(name string) watcherInterface.Interface {
	return WatcherExchangeRatio{
		watcherInterface.WatcherBase{
			Name_: name,
		},
	}
}

func (WatcherExchangeRatio) Call(ctx context.Context) watcherInterface.ResultInterface {
	ex, err := getSCExchangeRatio(ctx)
	return watcherInterface.ResultBase{
		Err:       err,
		IsNotify_: isSCBMarketOpen(utils.TimeNow()),
		Msg_:      fmt.Sprintf("Hi Boss! The buy ratio right now is %v, sell ratio is %v~", ex.BuyRatio, ex.SellRatio),
		Subject_:  "ResultExchangeRatio",
	}
}

type WatcherExchangeRatioLowerBuyRatio struct {
	watcherInterface.WatcherBase
	LowBoundRatio float64
}

func NewWatcherExchangeRatioLowerBuyRatio(name string, lowBoundRatio float64) watcherInterface.Interface {
	return WatcherExchangeRatioLowerBuyRatio{
		WatcherBase: watcherInterface.WatcherBase{
			Name_: name,
		},
		LowBoundRatio: lowBoundRatio,
	}
}

func (w WatcherExchangeRatioLowerBuyRatio) Call(ctx context.Context) watcherInterface.ResultInterface {
	ex, err := getSCExchangeRatio(ctx)
	isNotify := false
	if ex.BuyRatio < w.LowBoundRatio && isSCBMarketOpen(utils.TimeNow()) {
		isNotify = true
	}
	return watcherInterface.ResultBase{
		Err:       err,
		IsNotify_: isNotify,
		Msg_:      fmt.Sprintf("Hi Boss! The buy ratio right now is %v, lower than the bound %v~", ex.BuyRatio, w.LowBoundRatio),
		Subject_:  "ResultExchangeRatioLowerBuyRatio",
	}
}
