package hashx

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Sum(input ...[]byte) []byte {
	h := sha256.New()
	for _, v := range input {
		h.Write(v)
	}
	return h.Sum(nil)
}

func Sha256Hex(input ...[]byte) string {
	h := sha256.New()
	for _, v := range input {
		h.Write(v)
	}
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}
