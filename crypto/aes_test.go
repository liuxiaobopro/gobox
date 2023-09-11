package crypto

import (
	"testing"
)

func TestPwd_EnPwdCode(t *testing.T) {
	cr := Pwd{Key: []byte("1234567890123456")}

	str, _ := cr.EnPwdCode([]byte("123456"))

	t.Log(str)
}

func TestPwd_DePwdCode(t *testing.T) {
	cr := Pwd{Key: []byte("1234567890123456")}

	if str, err := cr.DePwdCode("VlRKR2MyUkhWbXRZTVRoaWVVSlFUbEZHTUROWU1uSXlNM05TUjNKTldHOVRUMWx0WWxaVGFIUXlXVDA9="); err != nil {
		t.Error(err)
	} else {
		t.Log(string(str))
	}
}
