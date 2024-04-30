package crypto

import "testing"

func TestMd5(t *testing.T) {
	md5str := Md5("123456")

	t.Log(md5str)
}
