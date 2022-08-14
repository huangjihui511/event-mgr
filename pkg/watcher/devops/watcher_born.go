package devops

import (
	"context"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/interfaces"
)

var (
	_ watcherInterface.Interface = &WatcherBorn{}
)

type WatcherBorn struct {
	isBorn bool
}

func NewWatcherBorn() watcherInterface.Interface {
	return &WatcherBorn{}
}

// Call implements watcherInterface.Interface
func (w *WatcherBorn) Call(ctx context.Context) watcherInterface.ResultInterface {
	isNotify := false
	if !w.isBorn {
		isNotify = true
	}
	w.isBorn = true
	return watcherInterface.ResultBase{
		IsNotify_: isNotify,
		Err:       nil,
		Msg_:      "Thanks you for giving me a new life",
		Subject_:  "system updated",
	}

}

// Name implements watcherInterface.Interface
func (WatcherBorn) Name() string {
	return "Born"
}
