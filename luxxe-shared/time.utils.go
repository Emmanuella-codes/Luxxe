package shared

import (
	"time"
)

func MinutesFromNow(minutes int, dayRequired ...time.Time) time.Time {
	baseTime := time.Now()
	if len(dayRequired) > 0 {
		baseTime = dayRequired[0]
	}
	return baseTime.Add(time.Duration(minutes) * time.Minute)
}

func GetCurrentYear() int {
	currentTime := time.Now().UTC()
	return currentTime.Year()
}
