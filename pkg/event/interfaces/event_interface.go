package interfaces

import (
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/watcher_interface"
)

type Interface interface {
	Chan() <-chan interface{}
	Watcher() watcherInterface.Interface
}
