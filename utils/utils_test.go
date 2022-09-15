package utils

import (
	"testing"
)

func TestGetLocalIPs(t *testing.T) {
	t.Log(GetLocalIPs())
}

func TestGetCurrentDir(t *testing.T) {
	t.Log(getCurrentDir())
}

func TestListDir(t *testing.T) {
	t.Log(listDir("."))
}
