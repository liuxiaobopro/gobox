package http

import (
	"github.com/gin-gonic/gin"
)

type GinHandle struct{}

// ShouldBind Get绑定到结构体
func (*GinHandle) ShouldBind(c *gin.Context, s interface{}) error {
	return c.ShouldBind(&s)
}

// ShouldBindJSON Post绑定到结构体
func (*GinHandle) ShouldBindJSON(c *gin.Context, s interface{}) error {
	return c.ShouldBindJSON(&s)
}

// Param 获取路由参数
func (*GinHandle) Param(c *gin.Context, key string) string {
	return c.Param(key)
}

// Query 获取get参数
func (*GinHandle) Query(c *gin.Context, key string) string {
	return c.Query(key)
}

// DefaultQuery 获取get参数，如果没有则返回默认值
func (*GinHandle) DefaultQuery(c *gin.Context, key string, defaultValue string) string {
	return c.DefaultQuery(key, defaultValue)
}
