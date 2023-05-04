package http

import (
	"mime/multipart"
	"net/http"

	respx "github.com/liuxiaobopro/gobox/resp"

	"github.com/gin-gonic/gin"
)

type GinHandle struct{}

// ShouldBind Get绑定到结构体
// 例如：
// type DemoGetReq struct {
// 	Id int `json:"id" form:"id"` // 必须有form标签
// }
// var r req.DemoGetReq
// /user?name=xxx => ShouldBind(c, &r)
func (*GinHandle) ShouldBind(c *gin.Context, s interface{}) error {
	return c.ShouldBind(s)
}

// ShouldBindJSON Post绑定到结构体
// 例如：
// type DemoPostReq struct {
// 	Id int `json:"id" form:"id"` // 必须有json标签,application/json
// }
// var r req.DemoPostReq
// /user => ShouldBindJSON(c, &r)
func (*GinHandle) ShouldBindJSON(c *gin.Context, s interface{}) error {
	return c.ShouldBindJSON(s)
}

// Param 获取路由参数
// 例如：/user/:name => Param(c, "name")
func (*GinHandle) Param(c *gin.Context, key string) string {
	return c.Param(key)
}

// Query 获取get参数
// 例如：/user?name=xxx => Query(c, "name")
func (*GinHandle) Query(c *gin.Context, key string) string {
	return c.Query(key)
}

// DefaultQuery 获取get参数，如果没有则返回默认值
// 例如：/user?name=xxx => DefaultQuery(c, "name", "default")
func (*GinHandle) DefaultQuery(c *gin.Context, key string, defaultValue string) string {
	return c.DefaultQuery(key, defaultValue)
}

// FormFile 获取上传文件
// 例如：FormFile(c, "file")
func (*GinHandle) FormFile(c *gin.Context, key string) (*multipart.FileHeader, error) {
	return c.FormFile(key)
}

// ReturnJSON 返回json
// 例如：ReturnJSON(c, http.StatusOK, respx.Succ(data))
func (*GinHandle) ReturnJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

// RetuenOk 返回成功json
// 例如：RetuenOk(c, data) => ReturnJSON(c, http.StatusOK, respx.Succ(data))
func (*GinHandle) RetuenOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, respx.Succ(data))
}

// ReturnErr 返回错误json
// 例如：ReturnErr(c, respx.ParamErrT.ToPt()) => ReturnJSON(c, http.StatusBadRequest, respx.ParamErrT.ToPt())
func (*GinHandle) ReturnErr(c *gin.Context, r *respx.Pt) {
	c.JSON(http.StatusBadRequest, r)
}
