package constants

import "time"

func RegularTimeEnter() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, time.Local)
}

func RegularTimeLeave() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, time.Local)
}
