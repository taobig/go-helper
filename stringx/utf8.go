package stringx

import (
	"strings"
	"unicode/utf8"
)

func IsUTF8(str string) bool {
	return utf8.ValidString(str)
}

func Len(str string) int {
	//if len(str) == 0 {
	//	return 0
	//}
	//r := []rune(str)
	//return len(r)
	return utf8.RuneCountInString(str)
}

// Substring returns a substring of str from start position to end position.
// If start is negative, it is treated as strLen + start where strLen is the length of str.
// If end is negative, it is treated as strLen + end where strLen is the length of str.
// If start is greater than or equal to end, a zero-length string is returned.
// If end is greater than the length of str, it is treated as the length of str.
func Substring(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rs := []rune(str)
	runeLen := len(rs)
	if runeLen == 0 {
		return ""
	}

	if start < 0 {
		start = runeLen + start
	}
	if start < 0 {
		start = 0
	}
	if start > runeLen-1 {
		return ""
	}

	end := runeLen
	if length < 0 {
		end = runeLen + length
	} else if length > 0 {
		end = start + length
	}

	if end < 0 || start >= end {
		return ""
	}
	if end > runeLen {
		end = runeLen
	}

	return string(rs[start:end])
}

func Cut(str string, start, end int) string {
	//runeSlice := []rune(str)
	//str = string(runeSlice[start:end])
	//return str
	return Substring(str, start, end-start)
}

func LowerFirstLetter(str string) string {
	return strings.ToLower(Substring(str, 0, 1)) + Substring(str, 1, len(str)) // len(str) >= len([]rune(str))
}

func UpperFirstLetter(str string) string {
	return strings.ToUpper(Substring(str, 0, 1)) + Substring(str, 1, len(str)) // len(str) >= len([]rune(str))
}
