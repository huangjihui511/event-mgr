package scb

import (
	"context"
	"fmt"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/watcher_interface"
)

var (
	_ watcherInterface.Interface       = ExchangeRatioWatcher{}
	_ watcherInterface.ResultInterface = ExchangeRatioResult{}
)

type ExchangeRatioWatcher struct {
	Name_ string
}

type ExchangeRatioResult struct {
	ExchangeRatio
	isNotify bool
	err      error
}

func (_ ExchangeRatioWatcher) Call(ctx context.Context) watcherInterface.ResultInterface {
	ex, err := getSCExchangeRatio(ctx)
	return ExchangeRatioResult{
		ExchangeRatio: ex,
		err:           err,
	}
}

func (e ExchangeRatioWatcher) Name() string {
	return e.Name_
}

func (e ExchangeRatioResult) Error() error {
	return e.err
}

func (e ExchangeRatioResult) Info() string {
	return fmt.Sprintf("%+v\n", e)
}

func (e ExchangeRatioResult) Subject() string {
	return "ExchangeRatioResult"
}

func (e ExchangeRatioResult) IsNotify() bool {
	return true
}
