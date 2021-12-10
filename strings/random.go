package strings

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

const (
	Digits            = 1 << iota // 1; 0b0001
	LowerCaseAlphabet             // 2; 0b0010
	UpperCaseAlphabet             // 4; 0b0100
)

func Random(length, numType int) (string, error) {
	if length <= 0 {
		return "", nil
	}
	arr := [62]string{}
	var index int = 0
	if numType&Digits == Digits {
		for i := 0; i <= 9; i++ {
			arr[index] = strconv.Itoa(i)
			index++
		}
	}
	if numType&LowerCaseAlphabet == LowerCaseAlphabet {
		for i := 'a'; i <= 'z'; i++ {
			arr[index] = string(i)
			index++
		}
	}
	if numType&UpperCaseAlphabet == UpperCaseAlphabet {
		for i := 'A'; i <= 'Z'; i++ {
			arr[index] = string(i)
			index++
		}
	}
	if index == 0 {
		return "", errors.New("invalid argument")
	}
	var str string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		randIndex := rand.Intn(index)
		//randIndex := fastrand.Uint32n(index) //github.com/valyala/fastrand
		str += arr[randIndex]
	}
	return str, nil
}
