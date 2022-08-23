package event

import (
	"context"
	"fmt"
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	"huangjihui511/event-mgr/pkg/event/timer"
	"huangjihui511/event-mgr/pkg/logs"
	"huangjihui511/event-mgr/pkg/notify"
	notifyInterface "huangjihui511/event-mgr/pkg/notify/interfaces"
	"huangjihui511/event-mgr/pkg/utils"
	"huangjihui511/event-mgr/pkg/watcher/devops"
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"huangjihui511/event-mgr/pkg/watcher/zsb"
	"time"
)

var (
	events       []eventInterface.Interface
	notifyEmail  notifyInterface.EmailInterface
	targetEmails = []string{
		"717655909@qq.com",
	}
)

func init() {
	notifyEmail = notify.EmailSender{
		EmailMeta: notifyInterface.EmailMetaQQ,
	}
}

func registerEvents() {
	events = []eventInterface.Interface{
		timer.NewTimer(time.Second, devops.NewWatcherBorn()),
		timer.NewTimer(time.Minute*30, scb.NewWatcherExchangeRatio("SCB Watcher")),
		timer.NewTimer(time.Minute, scb.NewWatcherExchangeRatioLowerBuyRatio("SCB Low Bound Watcher 6.8", 6.8)),
		timer.NewTimer(time.Minute*30, zsb.NewWatcherExchangeRatio("ZSB Watcher")),
		timer.NewTimer(time.Minute, zsb.NewWatcherExchangeRatioLowerBuyRatio("ZSB Low Bound Watcher 6.8", 6.8)),
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
		err := notifyEmail.Send(t, subject, r.Msg())
		if err != nil {
			logs.Logger.Errorf("send email to %v failed: %s", t, err)
		}
	}
}
