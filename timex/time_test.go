package timex

import (
	"testing"
	"time"
)

var cstZone = time.FixedZone("GMT", 8*3600) // UTC+08

func TestMillSecToTime(t *testing.T) {
	var ts int64 = 1640995200123 // 2022-01-01 00:00:00.123 +0000 UTC

	{
		actual := MillSecToTime(ts + 1).UTC().Format(time.RFC3339Nano)
		expected := "2022-01-01T00:00:00.124Z"
		if actual != expected {
			t.Errorf("actual:%s; expected %s", actual, expected)
		}
	}
	{
		actual := MillSecToTime(ts).In(cstZone).Format(time.RFC3339Nano)
		expected := "2022-01-01T08:00:00.123+08:00"
		if actual != expected {
			t.Errorf("actual:%s; expected %s", actual, expected)
		}
		//t.Log(actual)
	}

}

func TestMillSecToUtcTime(t *testing.T) {
	var ts int64 = 1640995200000 // 2022-01-01 00:00:00 +0000 UTC

	{
		actual := MillSecToUtcTime(ts).Format(time.RFC3339)
		expected := "2022-01-01T00:00:00Z"
		if actual != expected {
			t.Errorf("actual:%s; expected %s", actual, expected)
		}
		//t.Log(actual)
	}

}

func TestFormatTime(t *testing.T) {
	{
		var t1 time.Time = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
		actual := FormatTime(&t1, time.RFC3339)
		expected := "2022-01-01T00:00:00Z"
		if actual != expected {
			t.Errorf("actual:%s; expected %s", actual, expected)
		}
	}

	{
		var t1 *time.Time = nil
		actual := FormatTime(t1, time.RFC3339)
		expected := ""
		if actual != expected {
			t.Errorf("actual:%s; expected %s", actual, expected)
		}
	}

}

func isEqual(actual, expected time.Time) bool {
	//if actual.Year() != expected.Year() || actual.Month() != expected.Month() || actual.Day() != expected.Day() ||
	//	actual.Hour() != expected.Hour() || actual.Minute() != expected.Minute() ||
	//	actual.Second() != expected.Second() || actual.Nanosecond() != expected.Nanosecond() {
	return actual.UnixNano() == expected.UnixNano()
}

func TestStartOfDay(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.Local)
		actual := StartOfDay(start)
		expected := time.Date(2000, 1, 2, 0, 0, 0, 0, time.Local)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.UTC)
		actual := StartOfDay(start)
		expected := time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestTomorrow(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 1, 1, 1, 1, 111222333, time.Local)
		actual := StartOfTomorrow(start)
		expected := time.Date(2000, 1, 2, 0, 0, 0, 0, time.Local)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

	{
		start := time.Date(2000, 1, 1, 1, 1, 1, 111222333, time.UTC)
		actual := StartOfTomorrow(start)
		expected := time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestYesterday(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.Local)
		actual := StartOfYesterday(start)
		expected := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.UTC)
		actual := StartOfYesterday(start)
		expected := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestEndOfDay(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.Local)
		actual := EndOfDay(start)
		expected := time.Date(2000, 1, 2, 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.UTC)
		actual := EndOfDay(start)
		expected := time.Date(2000, 1, 2, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestEndOfTomorrow(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.Local)
		actual := EndOfTomorrow(start)
		expected := time.Date(2000, 1, 3, 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.UTC)
		actual := EndOfTomorrow(start)
		expected := time.Date(2000, 1, 3, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestEndOfYesterday(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.Local)
		actual := EndOfYesterday(start)
		expected := time.Date(2000, 1, 1, 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.UTC)
		actual := EndOfYesterday(start)
		expected := time.Date(2000, 1, 1, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestStartOfHour(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.Local)
		actual := StartOfHour(start)
		expected := time.Date(2000, 1, 2, 1, 0, 0, 0, time.Local)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.UTC)
		actual := StartOfHour(start)
		expected := time.Date(2000, 1, 2, 1, 0, 0, 0, time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestEndOfHour(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.Local)
		actual := EndOfHour(start)
		expected := time.Date(2000, 1, 2, 1, 59, 59, int(time.Second-time.Nanosecond), time.Local)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

	{
		start := time.Date(2000, 1, 2, 1, 1, 1, 111222333, time.UTC)
		actual := EndOfHour(start)
		expected := time.Date(2000, 1, 2, 1, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestStartOfMonth(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 10, 1, 1, 1, 111222333, time.UTC)
		actual := StartOfMonth(start)
		expected := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestEndOfMonth(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 1, 10, 1, 1, 1, 111222333, time.UTC)
		actual := EndOfMonth(start)
		expected := time.Date(2000, 1, 31, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

	{
		start := time.Date(2000, 2, 10, 1, 1, 1, 111222333, time.UTC)
		actual := EndOfMonth(start)
		expected := time.Date(2000, 2, 29, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestStartOfYear(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 2, 10, 1, 1, 1, 111222333, time.UTC)
		actual := StartOfYear(start)
		expected := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}

func TestEndOfYear(t *testing.T) {
	t.Parallel()

	{
		start := time.Date(2000, 2, 10, 1, 1, 1, 111222333, time.UTC)
		actual := EndOfYear(start)
		expected := time.Date(2000, 12, 31, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
		if !isEqual(actual, expected) {
			t.Logf("time is " + start.Format(time.RFC3339Nano))
			t.Errorf("actual:%s; expected %s", actual.Format(time.RFC3339Nano), expected.Format(time.RFC3339Nano))
		}
	}

}
