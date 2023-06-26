package log

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinConfig_Infof(t *testing.T) {
	logger := NewGin(
		WithMode(FileMode),
		WithLevel(InfoLevel),
		WithIsCloseColor(false),
	)
	c := &gin.Context{}
	logger.Infof(c, "hello %s", "world")
	logger.Errorf(c, "hello %s", "err")
}
