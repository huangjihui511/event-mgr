package scb

import (
	"context"
	"fmt"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/watcher_interface"
	"time"
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

func (_ WatcherExchangeRatio) Call(ctx context.Context) watcherInterface.ResultInterface {
	ex, err := getSCExchangeRatio(ctx)
	return ResultExchangeRatio{
		ExchangeRatio: ex,
		err:           err,
		isNotify:      isSCBMarketOpen(time.Now()),
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
