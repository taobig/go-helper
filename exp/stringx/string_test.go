package stringx

import (
	"testing"
)

func TestUpperFirstLetter(t *testing.T) {

	{
		source := "a"
		actual := UpperFirstLetter(source)
		expected := "A"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "aa"
		actual := UpperFirstLetter(source)
		expected := "Aa"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "hEllo"
		actual := UpperFirstLetter(source)
		expected := "HEllo"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "hello world"
		actual := UpperFirstLetter(source)
		expected := "Hello world"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "here comes o'brian"
		actual := UpperFirstLetter(source)
		expected := "Here comes o'brian"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	// 'Ğ'=>'ğ','Ü'=>'ü','Ş'=>'ş','İ'=>'i','Ö'=>'ö','Ç'=>'ç','I'=>'ı'
	{
		source := "ğüşiöçı"
		actual := UpperFirstLetter(source)
		expected := "Ğüşiöçı"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "你好，世界"
		actual := UpperFirstLetter(source)
		expected := "你好，世界"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "\taa"
		actual := UpperFirstLetter(source)
		expected := "\taa"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

}

func TestLowerFirstLetter(t *testing.T) {
	{
		source := "A"
		actual := LowerFirstLetter(source)
		expected := "a"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "Aa"
		actual := LowerFirstLetter(source)
		expected := "aa"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "HEllo"
		actual := LowerFirstLetter(source)
		expected := "hEllo"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "Hello World"
		actual := LowerFirstLetter(source)
		expected := "hello World"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "Here comes o'brian"
		actual := LowerFirstLetter(source)
		expected := "here comes o'brian"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	// 'Ğ'=>'ğ','Ü'=>'ü','Ş'=>'ş','İ'=>'i','Ö'=>'ö','Ç'=>'ç','I'=>'ı'
	{
		source := "ĞüşiöçI"
		actual := LowerFirstLetter(source)
		expected := "ğüşiöçI"
		if actual != expected {
			t.Errorf("LowerFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "你好，世界"
		actual := UpperFirstLetter(source)
		expected := "你好，世界"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

	{
		source := "\tAA"
		actual := UpperFirstLetter(source)
		expected := "\tAA"
		if actual != expected {
			t.Errorf("UpperFirstLetter(%v):%v; expected %v", source, actual, expected)
		}
	}

}
