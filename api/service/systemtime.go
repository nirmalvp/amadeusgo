package service

import "time"

type SystemTimeGetter struct{}

func (stg SystemTimeGetter) GetCurrentTime() time.Time {
	return time.Now()
}
