package time

import (
	"time"
)

// time.Now() returns the current local time.
func GetDateTimeWithNanosecondString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.999999999") //必须用这个特殊的数字2006-01-02 15:04:05  才能输入当前时间2014-11-23 14:24:03.1305627
}

// time.Now() returns the current local time.
func GetDateTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05") //必须用这个特殊的数字2006-01-02 15:04:05  才能输入当前时间2014-11-23 14:24:03.1305627
}

//layout  format:2006-01-02 15:04:05.999999999
func Parse(layout, date string) (time.Time, error) {
	if layout == "" {
		layout = "2006-01-02 15:04:05.999999999"
	}
	//t, err := time.Parse(layout, date)
	//time.Parse()的默认时区是UTC，time.Format()的时区默认是本地。
	t, err := time.ParseInLocation(layout, date, time.Local)
	if err != nil {
		return t, err
	}
	return t, nil
}
