package time

import (
	"github.com/taobig/go-helper/magic"
	"time"
)

// GetDateTimeWithNanosecondString
// Deprecated
func GetDateTimeWithNanosecondString(t time.Time) string {
	return t.Format(magic.SimpleDatetimeNanoLayout) //必须用这个特殊的数字2006-01-02 15:04:05  才能输入当前时间2014-11-23 14:24:03.1305627
}

// GetDateTimeString
// Deprecated
func GetDateTimeString(t time.Time) string {
	return t.Format(magic.SimpleDatetimeLayout) //必须用这个特殊的数字2006-01-02 15:04:05  才能输入当前时间2014-11-23 14:24:03
}

// Parse
// Deprecated
func Parse(layout, date string) (time.Time, error) {
	if layout == "" {
		layout = magic.SimpleDatetimeNanoLayout
	}
	//t, err := time.Parse(layout, date)
	//time.Parse()的默认时区是UTC，time.Format()的时区默认是本地。
	t, err := time.ParseInLocation(layout, date, time.Local)
	if err != nil {
		return t, err
	}
	return t, nil
}

// ParseInLocalLocation
// Deprecated
func ParseInLocalLocation(layout, date string) (time.Time, error) {
	//t, err := time.Parse(layout, date)
	//time.Parse()的默认时区是UTC，time.Format()的时区默认是本地。
	return time.ParseInLocation(layout, date, time.Local)
}
