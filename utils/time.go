package utils

import "time"

var (
	loc *time.Location
)

func init()  {
	loc, _ = time.LoadLocation("Asia/Shanghai")
}

func GetCurrentTime() string {
	//curTime := time.Now().Format("2006-01-02 15:04:05")
	curTime := time.Now().In(loc).Format("2006-01-02 15:04:05")

	return curTime
}
