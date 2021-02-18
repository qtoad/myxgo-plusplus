package util

import (
	"fmt"
	"testing"
)

func TestEncodeMD5(t *testing.T) {
	md5 := EncodeMd5("123456")
	fmt.Println(md5)
}
