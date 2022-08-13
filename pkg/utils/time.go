package utils

import "time"

var (
	Location *time.Location
)

func init() {
	Location, _ = time.LoadLocation("Asia/Shanghai")
}
