package interfaces

import (
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/interfaces"
)

//go:generate sh -c "mockgen --build_flags=--mod=mod huangjihui511/event-mgr/pkg/event/interfaces Interface > ./mock_interfaces/event_interface.go"

type Interface interface {
	Chan() <-chan interface{}
	Watcher() watcherInterface.Interface
}
