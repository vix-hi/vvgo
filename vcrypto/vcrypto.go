package goutils

import (
	"crypto/sha256"
	_"crypto/md5"
)


func Sha256Byte(sha256Str string) []byte {
	h := sha256.New()
	h.Write([]byte(sha256Str))
	return h.Sum(nil)
}

