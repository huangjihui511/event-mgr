package devops

import (
	"context"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/watcher_interface"
)

var (
	_ watcherInterface.Interface = WatcherBorn{}
)

type WatcherBorn struct {
}

// Call implements watcherInterface.Interface
func (WatcherBorn) Call(ctx context.Context) watcherInterface.ResultInterface {
	panic("unimplemented")
}

// Name implements watcherInterface.Interface
func (WatcherBorn) Name() string {
	panic("unimplemented")
}
