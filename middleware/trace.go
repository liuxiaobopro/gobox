package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/gobox/define"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取 Trace ID，如果不存在则生成一个新的 Trace ID
		traceId := c.GetString(define.TraceId)
		if traceId == "" {
			traceId = fmt.Sprintf("%d", time.Now().UnixNano())
		}

		// 将 Trace ID 存储到请求的上下文中
		c.Set(define.TraceId, traceId)

		// 继续处理请求
		c.Next()
	}
}
