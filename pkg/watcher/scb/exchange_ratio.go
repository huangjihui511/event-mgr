package scb

import (
	"context"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/watcher_interface"
)

var (
	_ watcherInterface.Interface       = ExchangeRatioWatcher{}
	_ watcherInterface.ResultInterface = ExchangeRatioResult{}
)

type ExchangeRatioWatcher struct {
}

type ExchangeRatioResult struct {
	ExchangeRatio
	err error
}

func (_ ExchangeRatioWatcher) Trigger(ctx context.Context) watcherInterface.ResultInterface {
	ex, err := getSCExchangeRatio(ctx)
	return ExchangeRatioResult{
		ExchangeRatio: ex,
		err:           err,
	}
}

func (e ExchangeRatioResult) Error() error {
	return e.err
}

func (e ExchangeRatioResult) String() string {
	panic("im")
}
