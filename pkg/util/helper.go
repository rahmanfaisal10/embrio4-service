package util

import "time"

func TimeNow() *time.Time {
	timer := time.Now()
	return &timer
}
