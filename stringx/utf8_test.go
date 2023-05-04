package stringx

import (
	"io/ioutil"
	"testing"
)

func TestSubstring(t *testing.T) {
	{
		str := ""
		actual := Substring(str, 0, 1)
		expected := ""
		if actual != expected {
			t.Errorf("Substring(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "你好，世界"
		actual := Substring(str, 0, 1)
		expected := "你"
		if actual != expected {
			t.Errorf("Substring(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "你好，世界"
		actual := Substring(str, 0, 0)
		expected := ""
		if actual != expected {
			t.Errorf("Substring(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "你好，世界"
		actual := Substring(str, 0, 100)
		expected := "你好，世界"
		if actual != expected {
			t.Errorf("Substring(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "你好，世界"
		actual := Substring(str, 0, -1)
		expected := "你好，世"
		if actual != expected {
			t.Errorf("Substring(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "你好，世界"
		actual := Substring(str, -3, -1)
		expected := "，世"
		if actual != expected {
			t.Errorf("Substring(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "你好，世界"
		actual := Substring(str, 0, 1)
		expected := str[:3]
		if actual != expected {
			t.Errorf("Substring(%v):%v; expected %v", str, actual, expected)
		}

		actual = Substring(str, 1, 3)
		expected = str[3:12]
		if actual != expected {
			t.Errorf("Substring(%v):%v; expected %v", str, actual, expected)
		}
	}
}

func TestCut(t *testing.T) {
	str := "你好，世界。hello world."

	var testCase = []struct {
		start    int
		end      int
		expected string
	}{
		{0, 1, "你"},
		{0, 2, "你好"},
		{0, 8, "你好，世界。he"},
		{0, 12, "你好，世界。hello "},
		{0, Len(str), str},
		{0, 100, str},
		{-14, 100, "界。hello world."},
	}

	for _, tt := range testCase {
		actual := Cut(str, tt.start, tt.end)
		if actual != tt.expected {
			t.Fatalf("actual:%s; expected %s\n", actual, tt.expected)
		}
	}
}

func TestIsUTF8(t *testing.T) {
	gbkBytes, err := ioutil.ReadFile("testdata/gbk.txt")
	if err != nil {
		t.Fatal(err)
	}
	gbkStr := string(gbkBytes)

	var testCase = []struct {
		input    string
		expected bool
	}{
		{"你好，世界。", true},
		{"你好，世界。he", true},
		{"你好，世界。hello ", true},
		{string(gbkStr), false},
	}

	for _, tt := range testCase {
		actual := IsUTF8(tt.input)
		if actual != tt.expected {
			t.Fatalf("actual:%v; expected %v\n", actual, tt.expected)
		}
	}
}

func TestLen(t *testing.T) {
	var testCase = []struct {
		input    string
		expected int
	}{
		{"你", 1},
		{"你好", 2},
		{"你好，世界。", 6},
		{"你好，世界。he", 8},
		{"你好，世界。hello ", 12},
	}

	for _, tt := range testCase {
		actual := Len(tt.input)
		if actual != tt.expected {
			t.Fatalf("actual:%d; expected %d\n", actual, tt.expected)
		}
	}
}

func TestUpperFirstLetter(t *testing.T) {
	{
		str := "hello"
		actual := UpperFirstLetter(str)
		expected := "Hello"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "hello world"
		actual := UpperFirstLetter(str)
		expected := "Hello world"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "你好"
		actual := UpperFirstLetter(str)
		expected := "你好"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		//德语Ä ä
		str := "ä"
		actual := UpperFirstLetter(str)
		expected := "Ä"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		//希腊字母 Δ δ
		str := "δ"
		actual := UpperFirstLetter(str)
		expected := "Δ"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

}

func TestLowerFirstLetter(t *testing.T) {
	{
		str := "Hello"
		actual := LowerFirstLetter(str)
		expected := "hello"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "HELLO world"
		actual := LowerFirstLetter(str)
		expected := "hELLO world"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		str := "你好"
		actual := LowerFirstLetter(str)
		expected := "你好"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		//德语Ä ä
		str := "Ä"
		actual := LowerFirstLetter(str)
		expected := "ä"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

	{
		//希腊字母 Δ δ
		str := "Δ"
		actual := LowerFirstLetter(str)
		expected := "δ"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", str, actual, expected)
		}
	}

}
