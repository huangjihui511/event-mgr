package pkg

import (
	"context"
	"huangjihui511/event-mgr/pkg/event"
	"huangjihui511/event-mgr/pkg/logs"
	"os"
	"os/signal"
	"time"
)

func StartService() {
	logs.Logger.Infof("Start service")
	ctx, cancel := context.WithCancel(context.Background())
	event.StartMgr(ctx)
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Kill)
	signal.Notify(exitChan, os.Interrupt)

	select {
	case <-exitChan:
		logs.Logger.Infof("Received stop signal, stopping smoothly")
		cancel()
		time.Sleep(time.Second)
	}
}
