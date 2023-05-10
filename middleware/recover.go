package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				_ = fmt.Errorf("[middleware recover] panic:%v", err)
				c.JSON(http.StatusInternalServerError, replyx.InternalErrT)
				c.Abort()
			}
		}()
		c.Next()
	}
}
