package event

import (
	"context"
	"fmt"
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	"huangjihui511/event-mgr/pkg/logs"
	"huangjihui511/event-mgr/pkg/notify"
	notifyInterface "huangjihui511/event-mgr/pkg/notify/interfaces"
)

func init() {
	notifyEmail = notify.EmailSender{
		EmailMeta: notifyInterface.EmailMetaQQ,
	}
}

func StartMgr(ctx context.Context) {
	registerEvents()
	startEvents(ctx)
}

func startEvents(ctx context.Context) {
	for _, e := range events {
		go func(ev eventInterface.Interface) {
			c := ev.Chan()
			do(ctx, ev)
			for {
				select {
				case <-c:
					do(ctx, ev)
				case <-ctx.Done():
					logs.Logger.Infof("Stop watcher %s", ev.Watcher().Name())
					return
				}
			}
		}(e)
	}
}

func do(ctx context.Context, ev eventInterface.Interface) {
	logs.Logger.Infof("Call watcher %s", ev.Watcher().Name())
	r := ev.Watcher().Call(ctx)
	if !r.IsNotify() {
		return
	}
	subject := fmt.Sprintf("%v: \"%v\"", ev.Watcher().Name(), r.Subject())
	for _, t := range targetEmails {
		err := notifyEmail.Send(t, subject, r.Msg())
		if err != nil {
			logs.Logger.Errorf("send email to %v failed: %s", t, err)
		}
	}
}
