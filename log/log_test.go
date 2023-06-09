package log

import (
	"testing"
)

func TestInfof(t *testing.T) {
	Infof("Hello %s", "World")
}

func TestErrorf(t *testing.T) {
	Errorf("Hello %s", "World")
}
