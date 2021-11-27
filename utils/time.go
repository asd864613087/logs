package utils

import (
	"time"
)

var (
	loc *time.Location
)

func init()  {
	// 空镜像没有tzdata
	// loc, _ = time.LoadLocation("Asia/Shanghai")
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	loc = time.FixedZone("CST", secondsEastOfUTC)
}

func GetCurrentTime() string {
	//curTime := time.Now().Format("2006-01-02 15:04:05")
	curTime := time.Now().In(loc).Format("2006-01-02 15:04:05")

	return curTime
}
