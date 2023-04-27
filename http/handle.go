package http

import (
	"mime/multipart"
	"net/http"

	respx "github.com/liuxiaobopro/gobox/resp"

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

// FormFile 获取上传文件
func (*GinHandle) FormFile(c *gin.Context, key string) (*multipart.FileHeader, error) {
	return c.FormFile(key)
}

// ReturnJSON 返回json
func (*GinHandle) ReturnJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

// RetuenOk 返回成功json
func (*GinHandle) RetuenOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, respx.Succ(data))
}

// ReturnErr 返回错误json
func (*GinHandle) ReturnErr(c *gin.Context, r respx.Pt) {
	c.JSON(http.StatusBadRequest, r)
}
