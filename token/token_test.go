package token

import (
	"testing"
)

var (
	sale       = "aaaaabbbbbcccccd"
	data       = `{"name":"liuxiaobo","age":18}`
	ecryptData = "ZdMoH5QScTz77IGVMzWbmApASH1G5lFcV+sSyqYgwxk="
)

func TestToken_GetToken(t *testing.T) {
	token := &Token{
		Sale: sale,
	}

	if str, err := token.GetToken(data); err != nil {
		t.Error(err)
	} else {
		t.Logf("GetToken success, token: %s", str)
	}
}

func TestToken_ParseToken(t *testing.T) {
	token := &Token{
		Sale: sale,
	}

	if str, err := token.ParseToken(ecryptData); err != nil {
		t.Error(err)
	} else {
		t.Logf("ParseToken success, data: %s", str)
	}
}
