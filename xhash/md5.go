package xhash

import (
	"github.com/taobig/go-helper/hashx"
)

// MD5 returns the md5 hash of the raw data
// Deprecated: use hashx.Md5Hex instead
func MD5(raw []byte) string {
	return hashx.Md5Hex(raw)
}
