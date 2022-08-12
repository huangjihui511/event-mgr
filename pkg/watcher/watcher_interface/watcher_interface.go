package watcherInterface

import "context"

type Interface interface {
	Call(ctx context.Context) ResultInterface
}

type ResultInterface interface {
	Error() error
	String() string
	IsNotify() bool
}
