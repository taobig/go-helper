package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocalIPs(t *testing.T) {
	t.Log(GetLocalIPs())
}

func TestGetCurrentDir(t *testing.T) {
	t.Log(getCurrentDir())
}

func TestListDir(t *testing.T) {
	assert.True(t, len(listDir(".")) > 0)
}
