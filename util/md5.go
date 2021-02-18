package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMd5 md5 encryption
func EncodeMd5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func Md5(value string) string {
	return EncodeMd5(value)
}
