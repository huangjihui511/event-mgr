package interfaces

import "context"

//go:generate sh -c "mockgen --build_flags=--mod=mod huangjihui511/event-mgr/pkg/watcher/interfaces Interface > ./mock_interfaces/watcher_interface.go"
//go:generate sh -c "mockgen --build_flags=--mod=mod huangjihui511/event-mgr/pkg/watcher/interfaces ResultInterface > ./mock_interfaces/result_interface.go"

var (
	_ ResultInterface = ResultBase{}
	_ Interface       = WatcherBase{}
)

type Interface interface {
	Call(ctx context.Context) ResultInterface
	Name() string
}

type ResultInterface interface {
	Error() error
	Msg() string
	IsNotify() bool
	Subject() string
}

type WatcherBase struct {
	Name_ string
}

// Call implements Interface
func (WatcherBase) Call(ctx context.Context) ResultInterface {
	panic("unimplemented")
}

// Name implements Interface
func (w WatcherBase) Name() string {
	return w.Name_
}

type ResultBase struct {
	IsNotify_ bool
	Err       error
	Msg_      string
	Subject_  string
}

// Error implements ResultInterface
func (r ResultBase) Error() error {
	return r.Err
}

// IsNotify implements ResultInterface
func (r ResultBase) IsNotify() bool {
	return r.IsNotify_
}

// Msg implements ResultInterface
func (r ResultBase) Msg() string {
	return r.Msg_
}

// Subject implements ResultInterface
func (r ResultBase) Subject() string {
	return r.Subject_
}
