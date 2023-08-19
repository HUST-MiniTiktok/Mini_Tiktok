package utils

import (
	"time"
)

func MillTimeStampToTime(timestamp int64) time.Time {
	sec := timestamp / 1000
	nsec := timestamp % 1000 * 1000000
	return time.Unix(sec, nsec)
}

func TimeToMillTimeStamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}