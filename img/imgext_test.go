package img

import "testing"

func TestGetImgExt(t *testing.T) {
	n := GetImgExt()
	if len(n) <= 0 {
		t.Fatalf("There must be atleast one extension provided")
	}
}
