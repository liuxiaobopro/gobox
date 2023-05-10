package log

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/gobox/define"
)

func TestInfof(t *testing.T) {
	// 创建一个默认的 gin 实例
	r := gin.Default()

	// 注册 GET 请求的路由处理函数
	r.GET("/", func(c *gin.Context) {
		c.Set(define.TRACE_ID, "1234567890")
		Infof(c, "Hello %s \n", "Gin")
		Infof(c, "%v \n", gin.H{
			"message": "Hello, world!",
		})
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// 启动 HTTP 服务器
	// r.Run(":8080")
}
