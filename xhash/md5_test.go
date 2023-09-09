package xhash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMD5(t *testing.T) {
	content := "hello world"
	hash := MD5([]byte(content))
	assert.Equal(t, "5eb63bbbe01eeed093cb22bb8f5acdc3", hash)
}
