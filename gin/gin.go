package gin

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	definex "github.com/liuxiaobopro/gobox/define"
)

func GetTraceId(c *gin.Context) string {
	return c.GetString(definex.TraceId)
}

func SetTraceId(c *gin.Context, traceId string) {
	c.Set(definex.TraceId, traceId)
}
func GetBody(c *gin.Context) string {
	var rMap map[string]interface{}
	_ = json.NewDecoder(c.Request.Body).Decode(&rMap)
	b, _ := json.Marshal(rMap)
	return string(b)
}
