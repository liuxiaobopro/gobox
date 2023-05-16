package token

import (
	cryptox "github.com/liuxiaobopro/gobox/crypto"
)

type Token struct {
	Sale string // 加密盐
}

func (t *Token) GetToken(data string) (string, error) {
	b, err := cryptox.AesEcrypt([]byte(data), []byte(t.Sale))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (t *Token) ParseToken(token string) (string, error) {
	b, err := cryptox.AesDeCrypt([]byte(token), []byte(t.Sale))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
