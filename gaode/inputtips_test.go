package gaode

import (
	"testing"
)

func TestNewInputtips(t *testing.T) {
	key := "xxx"

	tip := NewInputtips(key, WithKeywords("万达"))

	b, err := tip.Query()

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s", b)
}
