package hashx

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	content := "hello world"
	hash := Md5Sum([]byte(content))
	hexHash := hex.EncodeToString(hash[:])
	assert.Equal(t, "5eb63bbbe01eeed093cb22bb8f5acdc3", hexHash)
}

func TestMd5Hex(t *testing.T) {
	content := "hello world"
	hexHash := Md5Hex([]byte(content))
	assert.Equal(t, "5eb63bbbe01eeed093cb22bb8f5acdc3", hexHash)
}
