package ctx

import (
	"fmt"

	"github.com/gin-gonic/gin"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

type CtxInterface interface {
	Do()                                 // 业务逻辑
	ReturnJson(code int, data *replyx.T) // 返回 json
	Error(err error)                     // 错误处理
}

// 实现了 CtxInterface 接口的结构体
type Flow struct {
	ctx *gin.Context
}

func (f *Flow) Do() {
	panic("implement me")
}

func (f *Flow) ReturnJson(code int, data *replyx.T) {
	f.ctx.JSON(code, data)
}

func (f *Flow) Error(err error) {
	fmt.Printf("[Flow Error]: %v \n", err)
}
