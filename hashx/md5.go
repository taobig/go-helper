package hashx

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Sum(raw []byte) [16]byte {
	return md5.Sum(raw)
}

func Md5Hex(raw []byte) string {
	hash := md5.Sum(raw)
	hashHex := hex.EncodeToString(hash[:]) //"hello world" => 5eb63bbbe01eeed093cb22bb8f5acdc3
	return hashHex
}
