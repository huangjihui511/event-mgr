package scb

import (
	"context"
	"fmt"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/watcher_interface"
)

var (
	_ watcherInterface.Interface       = WatcherExchangeRatioLowerBuyRatio{}
	_ watcherInterface.ResultInterface = ResultExchangeRatioLowerBuyRatio{}
)

type WatcherExchangeRatioLowerBuyRatio struct {
	WatcherExchangeRatio
	LowBoundRatio float64
}

func (w WatcherExchangeRatioLowerBuyRatio) Call(ctx context.Context) watcherInterface.ResultInterface {
	ex, err := getSCExchangeRatio(ctx)
	isNotify := false
	if ex.BuyRatio < w.LowBoundRatio {
		isNotify = true
	}
	return ResultExchangeRatioLowerBuyRatio{
		ResultExchangeRatio: ResultExchangeRatio{
			ExchangeRatio: ex,
			err:           err,
			isNotify:      isNotify,
			msg:           fmt.Sprintf("Hi Boss! The sell ratio right now is %v, lower than the bound %v~", ex.BuyRatio, w.LowBoundRatio),
		},
	}
}

type ResultExchangeRatioLowerBuyRatio struct {
	ResultExchangeRatio
}

func (e ResultExchangeRatioLowerBuyRatio) Subject() string {
	return "ResultExchangeRatioLowerBuyRatio"
}
