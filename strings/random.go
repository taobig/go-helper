package strings

import (
	"github.com/taobig/go-helper/stringx"
)

const (
	// Deprecated:
	Digits = 1 << iota // 1; 0b0001
	// Deprecated:
	LowerCaseAlphabet // 2; 0b0010
	// Deprecated:
	UpperCaseAlphabet // 4; 0b0100
)

// Deprecated: As of 0.0.3, this function simply calls stringx.Random.
func Random(length, numType int) (string, error) {
	return stringx.Random(length, numType)
}
