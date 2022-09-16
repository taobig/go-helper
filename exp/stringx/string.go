package stringx

import (
	"unicode"
	"unicode/utf8"
)

func UpperFirstLetter(str string) string {
	if len(str) == 0 {
		return str
	}

	letter := str[0]
	if letter < utf8.RuneSelf {
		if 'a' <= letter && letter <= 'z' {
			letter -= 'a' - 'A'
		}
		return string(letter) + str[1:]
	}

	rs := []rune(str)
	return string(unicode.ToUpper(rs[0])) + string(rs[1:])
}

func LowerFirstLetter(str string) string {
	if len(str) == 0 {
		return str
	}

	letter := str[0]
	if letter < utf8.RuneSelf {
		if 'A' <= letter && letter <= 'Z' {
			letter += 'a' - 'A'
		}
		return string(letter) + str[1:]
	}

	rs := []rune(str)
	return string(unicode.ToLower(rs[0])) + string(rs[1:])
}
