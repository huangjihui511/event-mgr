package event

import (
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	"sync"
	"time"
)

type Dashboard struct {
	Items []DashboardItem
	Time  time.Time
	sync.RWMutex
}

type DashboardItem struct {
	Name         string
	Msg          string
	IsNotify     bool
	LastCallAt   time.Time
	LastNotifyAt time.Time
	Err          string
}

var DashboardData Dashboard

func initDashboard(events []eventInterface.Interface) {
	DashboardData.Items = make([]DashboardItem, len(events))
}
