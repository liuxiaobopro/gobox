package gin

import (
	"github.com/gin-gonic/gin"
	definex "github.com/liuxiaobopro/gobox/define"
)

func GetTraceId(c *gin.Context) string {
	return c.GetString(definex.TraceId)
}

func SetTraceId(c *gin.Context, traceId string) {
	c.Set(definex.TraceId, traceId)
}
