package zsb

import (
	"context"
	"fmt"
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
	ex, err := getZSExchangeRatio(ctx)
	return watcherInterface.ResultBase{
		Err:       err,
		IsNotify_: true,
		Msg_:      fmt.Sprintf("Hi Boss! The ZSB buy ratio right now is %v, sell ratio is %v~", ex.USDBuy, ex.USDSell),
		Subject_:  "get new ratio",
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
	ex, err := getZSExchangeRatio(ctx)
	isNotify := false
	msg := fmt.Sprintf("Hi Boss! The ZSB buy ratio right now is %v, higher than the bound %v~", ex.USDBuy, w.LowBoundRatio)
	if ex.USDBuy < w.LowBoundRatio {
		isNotify = true
		msg = fmt.Sprintf("Hi Boss! The ZSB buy ratio right now is %v, lower than the bound %v~", ex.USDBuy, w.LowBoundRatio)
	}
	return watcherInterface.ResultBase{
		Err:       err,
		IsNotify_: isNotify,
		Msg_:      msg,
		Subject_:  "buy ratio changed",
	}
}
