package utils

import "time"

var (
	Location *time.Location
	TimeNow  = time.Now
)

func init() {
	Location, _ = time.LoadLocation("Asia/Shanghai")
}
