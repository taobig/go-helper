package converter

import "strconv"

func StrToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func StrToUint(str string) (uint, error) {
	i64, err := strconv.ParseInt(str, 10, 64)
	return uint(i64), err
}

func StrToInt32(str string) (int32, error) {
	i64, err := strconv.ParseInt(str, 10, 64)
	return int32(i64), err
}

func StrToUint32(str string) (uint32, error) {
	i64, err := strconv.ParseInt(str, 10, 64)
	return uint32(i64), err
}

func StrToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func StrToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func UintToStr(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

func Int32ToStr(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Uint32ToStr(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Uint64ToStr(i uint64) string {
	return strconv.FormatUint(i, 10)
}
