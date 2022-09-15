package timex

import (
	"github.com/taobig/go-helper/magic"
	"testing"
	"time"
)

var cstZone = time.FixedZone("GMT", 8*3600) // UTC+08

func TestGetDateTime(t *testing.T) {
	// "2022-01-01T08:00:00Z" (UTC)
	// "2022-01-01T08:00:00+08:00" (UTC+8)
	t.Log(time.Now().Format(time.RFC3339))
	// "2006-01-02T15:04:05.999999999Z" (UTC)
	// "2006-01-02T15:04:05.999999999+08:00" (UTC+8)
	t.Log(time.Now().Format(time.RFC3339Nano))

	{
		//s := time.Now().In(time.UTC).Format("2006-01-02 00:00:00")
		s := time.Now().In(time.UTC).Format(magic.SimpleDateLayout)
		t.Log(s)
	}
	{
		s := time.Now().In(time.UTC).Format(magic.SimpleDatetimeLayout)
		t.Log(s)
	}
	{
		s := time.Now().In(time.UTC).Format(magic.SimpleDatetimeNanoLayout)
		t.Log(s)
	}

}

// time.Parse() Usage
func TestParse(t *testing.T) {
	t.Parallel()

	{
		_, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
		if err != nil {
			t.Error(err)
		}
	}

	{
		_, err := time.Parse(time.RFC3339, "2022-01-01T01:01:01+08:00")
		if err != nil {
			t.Error(err)
		}

		_, err = time.Parse(time.RFC3339, "2022-01-01T01:01:01Z")
		if err != nil {
			t.Error(err)
		}
	}

	{
		var date string = "2020-01-01 00:00:00.999999"
		_, err := time.Parse("", date)
		if err != nil {
			// `layout` param can't be ""
			//t.Errorf("error:%v", err)
		} else {
			t.Errorf("expected error, actual:nil")
		}
	}

	{
		date := "2020-01-01 00:00:00.999999"
		_, err := time.Parse(time.RFC3339Nano, date)
		if err == nil {
			t.Errorf("time.Parse(%v, %v) failed; expected error", time.RFC3339Nano, date)
		}
	}

}

// time.ParseInLocation() Usage
func TestParseInLocalLocation(t *testing.T) {
	t.Parallel()

	_, err := time.ParseInLocation(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015", time.Local)
	if err != nil {
		t.Error(err)
	}

	date := "2020-01-01 00:00:00.999999"
	_, err = time.ParseInLocation("", date, time.Local)
	if err != nil {
		// `layout` param can't be ""
		// t.Logf("expected error:%+v", err)
	} else {
		t.Errorf("expected error, actual:nil")
	}

	date = "2020-01-01T00:00:00.999999Z"
	_, err = time.ParseInLocation(time.RFC3339Nano, date, time.UTC)
	if err != nil {
		t.Errorf("Parse(time.RFC3339Nano, %v) failed; expected error", date)
		t.Error(err)
	}
}

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
