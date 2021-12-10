package strings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRandomStr(t *testing.T) {
	t.Parallel()

	actual, _ := Random(10, Digits)
	fmt.Println(actual)
	match, _ := regexp.MatchString("^[0-9]{10}$", actual)
	if !match {
		t.Errorf("Random():%s", actual)
	}
	actual, _ = Random(10, LowerCaseAlphabet)
	fmt.Println(actual)
	match, _ = regexp.MatchString("^[a-z]{10}$", actual)
	if !match {
		t.Errorf("Random():%s", actual)
	}
	actual, _ = Random(10, UpperCaseAlphabet)
	fmt.Println(actual)
	match, _ = regexp.MatchString("^[A-Z]{10}$", actual)
	if !match {
		t.Errorf("Random():%s", actual)
	}

	actual, _ = Random(10, Digits|LowerCaseAlphabet)
	fmt.Println(actual)
	match, _ = regexp.MatchString("^[0-9a-z]{10}$", actual)
	if !match {
		t.Errorf("Random():%s", actual)
	}
	actual, _ = Random(10, Digits|UpperCaseAlphabet)
	fmt.Println(actual)
	match, _ = regexp.MatchString("^[0-9A-Z]{10}$", actual)
	if !match {
		t.Errorf("Random():%s", actual)
	}
	actual, _ = Random(10, LowerCaseAlphabet|UpperCaseAlphabet)
	fmt.Println(actual)
	match, _ = regexp.MatchString("^[a-zA-Z]{10}$", actual)
	if !match {
		t.Errorf("Random():%s", actual)
	}
	actual, _ = Random(10, Digits|LowerCaseAlphabet|UpperCaseAlphabet)
	fmt.Println(actual)
	match, _ = regexp.MatchString("^[0-9a-zA-Z]{10}$", actual)
	if !match {
		t.Errorf("Random():%s", actual)
	}
	actual, _ = Random(52, Digits|LowerCaseAlphabet|UpperCaseAlphabet)
	fmt.Println(actual)
	match, _ = regexp.MatchString("^[0-9a-zA-Z]{52}$", actual)
	if !match {
		t.Errorf("Random():%s", actual)
	}
	actual, _ = Random(53, Digits|LowerCaseAlphabet|UpperCaseAlphabet)
	fmt.Println(actual)
	match, _ = regexp.MatchString("^[0-9a-zA-Z]{53}$", actual)
	if !match {
		t.Errorf("Random():%s", actual)
	}
}
