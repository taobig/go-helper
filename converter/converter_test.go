package converter

import (
	"math"
	"testing"
)

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestStrToNum(t *testing.T) {
	func() {
		actual, err := StrToInt("1")
		expected := 1
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}
	}()

	func() {
		actual, err := StrToUint("1")
		var expected uint = 1
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}
	}()

	func() {
		actual, err := StrToInt32("1")
		var expected int32 = 1
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}

		actual, err = StrToInt32("2147483647")
		expected = math.MaxInt32
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}
	}()

	func() {
		actual, err := StrToUint32("1")
		var expected uint32 = 1
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}

		actual, err = StrToUint32("4294967295")
		expected = math.MaxUint32
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}
	}()

	func() {
		actual, err := StrToInt64("1")
		var expected int64 = 1
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}

		actual, err = StrToInt64("9223372036854775807")
		expected = math.MaxInt64
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}
	}()

	func() {
		actual, err := StrToUint64("1")
		var expected uint64 = 1
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}

		actual, err = StrToUint64("18446744073709551615")
		expected = math.MaxUint64
		checkError(t, err)
		if actual != expected {
			t.Errorf("actual:%d; expected %d", actual, expected)
		}
	}()
}

func TestNumToStr(t *testing.T) {
	actual, expected := IntToStr(1), "1"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}

	actual, expected = UintToStr(1), "1"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}

	actual, expected = Int32ToStr(1), "1"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}

	actual, expected = Int32ToStr(math.MaxInt32), "2147483647"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}

	actual, expected = Uint32ToStr(1), "1"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}

	actual, expected = Uint32ToStr(math.MaxUint32), "4294967295"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}

	actual, expected = Int64ToStr(1), "1"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}

	actual, expected = Int64ToStr(math.MaxInt64), "9223372036854775807"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}

	actual, expected = Uint64ToStr(1), "1"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}

	actual, expected = Uint64ToStr(math.MaxUint64), "18446744073709551615"
	if actual != expected {
		t.Errorf("actual:%s; expected %s", actual, expected)
	}
}
