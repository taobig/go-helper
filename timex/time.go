package timex

import (
	"time"
)

func MillSecToTime(ts int64) time.Time {
	second := ts / 1000
	millSecond := ts % 1000
	return time.Unix(second, millSecond*1e6)
}

func MillSecToUtcTime(ts int64) time.Time {
	return MillSecToTime(ts).UTC() //or return MillSecToTime(ts).In(time.UTC)
}
