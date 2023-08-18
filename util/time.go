package util

import (
	"time"
)

func MillTimeStampToTime(timestamp int64) time.Time {
	sec := timestamp / 1000
	nsec := timestamp % 1000 * 1000000
	return time.Unix(sec, nsec)
}