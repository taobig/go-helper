package converter

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestStrToNum(t *testing.T) {
	{
		actual, err := StrToInt("1")
		expected := 1
		checkError(t, err)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	{
		actual, err := StrToUint("1")
		var expected uint = 1
		checkError(t, err)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	{
		actual, err := StrToInt32("1")
		var expected int32 = 1
		checkError(t, err)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	{
		actual, err := StrToInt32("2147483647")
		expected := math.MaxInt32
		checkError(t, err)
		assert.IsType(t, int32(0), actual, "they should be equal")
		assert.Equal(t, int64(expected), int64(actual), "they should be equal")
	}

	{
		actual, err := StrToUint32("1")
		var expected uint32 = 1
		checkError(t, err)
		assert.IsType(t, uint32(0), actual, "they should be equal")
		assert.Equal(t, expected, actual, "they should be equal")
	}
	{
		actual, err := StrToUint32("4294967295")
		expected := math.MaxUint32
		checkError(t, err)
		assert.IsType(t, uint32(0), actual, "they should be equal")
		assert.Equal(t, int64(expected), int64(actual), "they should be equal")
	}

	{
		actual, err := StrToInt64("1")
		var expected int64 = 1
		checkError(t, err)
		assert.IsType(t, int64(0), actual, "they should be equal")
		assert.Equal(t, int64(expected), int64(actual), "they should be equal")
	}
	{
		actual, err := StrToInt64("9223372036854775807")
		expected := math.MaxInt64
		checkError(t, err)
		assert.IsType(t, int64(0), actual, "they should be equal")
		assert.Equal(t, int64(expected), int64(actual), "they should be equal")
	}

	{
		actual, err := StrToUint64("1")
		var expected uint64 = 1
		checkError(t, err)
		assert.IsType(t, uint64(0), actual, "they should be equal")
		assert.Equal(t, uint64(expected), uint64(actual), "they should be equal")
	}
	{
		actual, err := StrToUint64("18446744073709551615")
		var expected uint64 = math.MaxUint64
		checkError(t, err)
		assert.IsType(t, uint64(0), actual, "they should be equal")
		assert.Equal(t, uint64(expected), uint64(actual), "they should be equal")
	}
}

func TestNumToStr(t *testing.T) {
	actual, expected := IntToStr(1), "1"
	assert.Equal(t, expected, actual, "they should be equal")

	actual, expected = UintToStr(1), "1"
	assert.Equal(t, expected, actual, "they should be equal")

	actual, expected = Int32ToStr(1), "1"
	assert.Equal(t, expected, actual, "they should be equal")

	actual, expected = Int32ToStr(math.MaxInt32), "2147483647"
	assert.Equal(t, expected, actual, "they should be equal")

	actual, expected = Uint32ToStr(1), "1"
	assert.Equal(t, expected, actual, "they should be equal")

	actual, expected = Uint32ToStr(math.MaxUint32), "4294967295"
	assert.Equal(t, expected, actual, "they should be equal")

	actual, expected = Int64ToStr(1), "1"
	assert.Equal(t, expected, actual, "they should be equal")

	actual, expected = Int64ToStr(math.MaxInt64), "9223372036854775807"
	assert.Equal(t, expected, actual, "they should be equal")

	actual, expected = Uint64ToStr(1), "1"
	assert.Equal(t, expected, actual, "they should be equal")

	actual, expected = Uint64ToStr(math.MaxUint64), "18446744073709551615"
	assert.Equal(t, expected, actual, "they should be equal")
}
