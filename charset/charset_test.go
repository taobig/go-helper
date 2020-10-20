package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	s := "GBK 与 UTF-8 编码转换测试"
	gbk, err := UTF82GBK([]byte(s))
	if err != nil {
		t.Errorf("utf-8 to gbk error:%v", err.Error())
	} else {
		//t.Log(string(gbk))
		//ioutil.WriteFile("test.gbk.txt", gbk, os.ModePerm)
		//ioutil.WriteFile("test.utf8.txt", []byte(s), os.ModePerm)

		_, err := UTF82GBK(gbk)
		if err != nil {
			//t.Logf("good job")
		} else {
			t.Fatalf("error: %v", gbk)
		}

		utf8, err := GBK2UTF8(gbk)
		if err != nil {
			t.Errorf("gbk to utf-8 error:%v", err.Error())
		} else {
			if len(utf8) != len(s) {
				t.Errorf("gbk to utf-8 error: length is incorrect ")
			}
		}

		_, err = GBK2UTF8([]byte(s))
		if err == nil {
			t.Errorf("gbk to utf-8 error: input params is a utf-8 string")
		}

	}
}
