package event

import (
	"context"
	"fmt"
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	"huangjihui511/event-mgr/pkg/logs"
	"huangjihui511/event-mgr/pkg/notify"
	notifyInterface "huangjihui511/event-mgr/pkg/notify/interfaces"
	"huangjihui511/event-mgr/pkg/utils"
)

var (
	events      []eventInterface.Interface
	notifyEmail notifyInterface.EmailInterface
)

func init() {
	notifyEmail = notify.EmailSender{
		EmailMeta: notifyInterface.EmailMetaQQ,
	}
}

func StartMgr(ctx context.Context) {
	registerEvents()
	startEvents(ctx, events)
}

func startEvents(ctx context.Context, events []eventInterface.Interface) {
	initDashboard(events)
	for i, e := range events {
		go func(ev eventInterface.Interface, index int) {
			c := ev.Chan(ctx)
			do(ctx, ev, index)
			for {
				select {
				case <-c:
					do(ctx, ev, index)
				case <-ctx.Done():
					logs.Logger.Infof("Stop watcher %s", ev.Watcher().Name())
					return
				}
			}
		}(e, i)
	}
}

func do(ctx context.Context, ev eventInterface.Interface, index int) {
	logs.Logger.Infof("Call watcher %s", ev.Watcher().Name())
	r := ev.Watcher().Call(ctx)
	DashboardData.Lock()
	defer DashboardData.Unlock()
	DashboardData.Items[index] = DashboardItem{
		Name:       ev.Watcher().Name(),
		Msg:        r.Msg(),
		IsNotify:   r.IsNotify(),
		LastCallAt: utils.TimeNow(),
	}
	if r.Error() != nil {
		DashboardData.Items[index].Err = r.Error().Error()
	}
	if !r.IsNotify() {
		return
	}
	DashboardData.Items[index].LastNotifyAt = utils.TimeNow()
	subject := fmt.Sprintf("%v: \"%v\"", ev.Watcher().Name(), r.Subject())
	for _, t := range targetEmails {
		ri := RouterItem{
			EventName: ev.Watcher().Name(),
			Email:     t,
		}
		if !RouterItems[ri] {
			continue
		}
		err := notifyEmail.Send(t, subject, r.Msg())
		if err != nil {
			logs.Logger.Errorf("send email to %v failed: %s", t, err)
		}
	}
}
