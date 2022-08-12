package watcherInterface

import "context"

type Interface interface {
	Trigger(ctx context.Context) ResultInterface
}

type ResultInterface interface {
	Error() error
	String() string
}
