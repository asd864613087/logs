package utils

import "time"

func GetCurrentTime() string {
	curTime := time.Now().Format("2006-01-02 15:04:05")
	return curTime
}
