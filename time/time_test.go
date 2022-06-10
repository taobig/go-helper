package time

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDateTime(t *testing.T) {
	fmt.Println(time.Now().Format(time.RFC3339))
}

func TestGetDateTimeWithNanosecondString(t *testing.T) {
	t.Parallel()

	fmt.Println(GetDateTimeWithNanosecondString(time.Now()))
}

func TestGetDateTimeString(t *testing.T) {
	t.Parallel()

	fmt.Println(GetDateTimeString(time.Now()))
}

func TestParse(t *testing.T) {
	t.Parallel()

	_, err := Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(Parse("", "2020-01-01 00:00:00"))
	//fmt.Println(Parse("", "2020-01-01 00:00:00.999999"))
	var date string = "2020-01-01 00:00:00.999999"
	t1, err := Parse("", date)
	if err != nil {
		t.Errorf("error:%v", err)
	} else {
		if t1.Year() != 2020 {
			t.Errorf("Parse(%v):%v; expected %v", date, 2020, 2020)
		}
	}

	date = "2020-01-01 00:00:00.999999"
	_, err = Parse(time.RFC3339Nano, date)
	if err == nil {
		t.Errorf("Parse(%v, %v) failed; expected error", time.RFC3339Nano, date)
	}

}

func TestParseInLocalLocation(t *testing.T) {
	t.Parallel()

	_, err := ParseInLocalLocation(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil {
		t.Error(err)
	}

	date := "2020-01-01 00:00:00.999999"
	_, err = ParseInLocalLocation("", date)
	if err != nil {
		t.Logf("expected error:%+v", err)
	} else {
		t.Errorf("expected error, actual:nil")
	}

	date = "2020-01-01 00:00:00.999999"
	_, err = ParseInLocalLocation(time.RFC3339Nano, date)
	if err == nil {
		t.Errorf("Parse(%v, %v) failed; expected error", time.RFC3339Nano, date)
	}
}
