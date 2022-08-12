package watcherInterface

import "context"

type Interface interface {
	Call(ctx context.Context) ResultInterface
	Name() string
}

type ResultInterface interface {
	Error() error
	Info() string
	IsNotify() bool
	Subject() string
}
