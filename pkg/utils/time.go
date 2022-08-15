package utils

import (
	"huangjihui511/event-mgr/pkg/logs"
	"time"
)

var (
	Location *time.Location
	TimeNow  = func() time.Time {
		return time.Now().In(Location)
	}
)

func init() {
	Location = time.Now().Location()
	logs.Logger.Infof("origin location is %v", Location.String())
	Location, _ = time.LoadLocation("Asia/Shanghai")
	time.Now().Local()
}
