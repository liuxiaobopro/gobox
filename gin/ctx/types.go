package ctx

import (
	"fmt"

	"github.com/gin-gonic/gin"
	definex "github.com/liuxiaobopro/gobox/define"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

const (
	authorCtxKey = "Author" // 作者上下文键
)

type ICtx interface {
	GetAuthor() string                   // 获取作者
	Do()                                 // 业务逻辑
	ReturnJson(code int, data *replyx.T) // 返回 json
	Error(err error)                     // 错误处理

	setCtx(ctx *gin.Context)
}

// 实现了 CtxInterface 接口的结构体
type Flow struct {
	ctx *gin.Context
}

func (f *Flow) GetAuthor() string {
	return f.ctx.GetString(authorCtxKey)
}

func (f *Flow) Do() {
	panic("implement func Do()")
}

func (f *Flow) ReturnJson(code int, data *replyx.T) {
	f.ctx.JSON(code, data)
}

func (f *Flow) Error(err error) {
	fmt.Printf("[Flow Error] %v \n", err)
}

func (f *Flow) setCtx(ctx *gin.Context) {
	f.ctx = ctx
	f.ctx.Set(authorCtxKey, definex.Author)
}
