package log

import (
	"context"
	"testing"
)

func TestInfof(t *testing.T) {
	Infof(context.Background(), "Hello %s", "World")
}

func TestErrorf(t *testing.T) {
	Errorf(context.Background(), "Hello %s", "World")
}
