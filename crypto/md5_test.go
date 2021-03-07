package crypto

import (
	"fmt"
	"testing"
)

func TestEncodeMD5(t *testing.T) {
	md5 := Md5("123456")
	fmt.Println(md5)
}
