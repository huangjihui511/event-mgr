package scb

import (
	"context"
	"fmt"
	"huangjihui511/event-mgr/pkg/utils"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/watcher_interface"
)

var (
	_ watcherInterface.Interface       = WatcherExchangeRatio{}
	_ watcherInterface.ResultInterface = ResultExchangeRatio{}
)

type WatcherExchangeRatio struct {
	Name_ string
}

type ResultExchangeRatio struct {
	ExchangeRatio
	isNotify bool
	err      error
	msg      string
}

func (WatcherExchangeRatio) Call(ctx context.Context) watcherInterface.ResultInterface {
	ex, err := getSCExchangeRatio(ctx)
	return ResultExchangeRatio{
		ExchangeRatio: ex,
		err:           err,
		isNotify:      isSCBMarketOpen(utils.TimeNow()),
		msg:           fmt.Sprintf("Hi Boss! The buy ratio right now is %v, sell ratio is %v~", ex.BuyRatio, ex.SellRatio),
	}
}

func (e WatcherExchangeRatio) Name() string {
	return e.Name_
}

func (e ResultExchangeRatio) Error() error {
	return e.err
}

func (e ResultExchangeRatio) Msg() string {
	return e.msg
}

func (e ResultExchangeRatio) Subject() string {
	return "ResultExchangeRatio"
}

func (e ResultExchangeRatio) IsNotify() bool {
	return e.isNotify
}
