package log

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/gobox/middleware"
)

func TestGinConfig_Infof(t *testing.T) {
	logger := NewGin(WithMode(FileMode), WithLevel(InfoLevel))

	r := gin.Default()
	r.Use(middleware.Trace())

	r.GET("/", func(c *gin.Context) {
		logger.Infof(c, "hello %s", "world")
		logger.Errorf(c, "hello %s", "err")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	// r.Run(":8080")
}
