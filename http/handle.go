package http

import (
	"net/http"

	respx "github.com/liuxiaobopro/gobox/resp"

	"github.com/gin-gonic/gin"
)

type GinHandle struct {
	c *gin.Context
}

// ShouldBind Get绑定到结构体
func (h *GinHandle) ShouldBind(s interface{}) error {
	return h.c.ShouldBind(&s)
}

// ShouldBindJSON Post绑定到结构体
func (h *GinHandle) ShouldBindJSON(s interface{}) error {
	return h.c.ShouldBindJSON(&s)
}

// Param 获取路由参数
func (h *GinHandle) Param(key string) string {
	return h.c.Param(key)
}

// Query 获取get参数
func (h *GinHandle) Query(key string) string {
	return h.c.Query(key)
}

// DefaultQuery 获取get参数，如果没有则返回默认值
func (h *GinHandle) DefaultQuery(key string, defaultValue string) string {
	return h.c.DefaultQuery(key, defaultValue)
}

// ReturnJSON 返回json
func (h *GinHandle) ReturnJSON(code int, data interface{}) {
	h.c.JSON(code, data)
}

// RetuenOk 返回成功json
func (h *GinHandle) RetuenOk(data interface{}) {
	h.c.JSON(http.StatusOK, respx.Succ(data))
}

// ReturnErr 返回错误json
func (h *GinHandle) ReturnErr(r respx.T) {
	h.c.JSON(http.StatusBadRequest, r)
}
