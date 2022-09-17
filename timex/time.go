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

func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	//return StartOfTomorrow(t).Add(-time.Nanosecond)
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

func StartOfTomorrow(t time.Time) time.Time {
	return StartOfDay(t.AddDate(0, 0, 1))
}

func EndOfTomorrow(t time.Time) time.Time {
	return EndOfDay(t.AddDate(0, 0, 1))
}

func StartOfYesterday(t time.Time) time.Time {
	return StartOfDay(t.AddDate(0, 0, -1))
}

func EndOfYesterday(t time.Time) time.Time {
	return EndOfDay(t.AddDate(0, 0, -1))
}

func StartOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

func StartOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t.AddDate(0, 1, 0)).Add(-time.Nanosecond)
}

func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

func EndOfYear(t time.Time) time.Time {
	return StartOfYear(t.AddDate(1, 0, 0)).Add(-time.Nanosecond)
}
