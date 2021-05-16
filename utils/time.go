package utils

import "time"

func GetCurrentTime() string {
	curTime := time.Now().String()
	return curTime
}
