package crypto

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	userPassword := "123111111111456"
	passwordbyte, err := GeneratePassword(userPassword)
	if err != nil {
		t.Errorf("GeneratePassword() error = %v", err)
	}

	t.Logf("password: %s \n", string(passwordbyte))

	if ok, err := ValidatePassword(userPassword, string(passwordbyte)); !ok {
		t.Errorf("ValidatePassword() error = %v", err)
	}
}
