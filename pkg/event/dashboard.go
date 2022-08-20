package event

import (
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
}

var DashboardData Dashboard
