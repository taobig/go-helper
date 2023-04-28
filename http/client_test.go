package http

import (
	"os"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	t.Parallel()

	{
		url := "https://github.com/taobig/go-helper/workflows/Go/badge.svg"
		path, err := DownloadFile(url, ".")
		if err != nil {
			t.Errorf("DownloadFile(%v) error:%v", url, err.Error())
		} else {
			t.Log(path)
		}

		path, err = DownloadFile(url, ".")
		if err != nil {
			t.Errorf("DownloadFile(%v) error:%v", url, err.Error())
		} else {
			t.Log(path)
			os.Remove(path)
		}
	}

}
