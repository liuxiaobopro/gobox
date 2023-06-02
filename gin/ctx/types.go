package ctx

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	definex "github.com/liuxiaobopro/gobox/define"
	logx "github.com/liuxiaobopro/gobox/log"
	replyx "github.com/liuxiaobopro/gobox/reply"
)

const (
	authorCtxKey = "Author" // 作者上下文键
)

type ICtx interface {
	GetAuthor() string                              // 获取作者
	ReturnJson(data *replyx.T)                      // 返回 json
	SetHttpCode(code int)                           // 设置 http code
	PrintErrorf(format string, args ...interface{}) // 打印错误日志
	PrintInfof(format string, args ...interface{})  // 打印信息日志

	Handle()   // 业务逻辑-控制器句柄
	Validate() // 业务逻辑-参数校验
	Logic()    // 业务逻辑-业务逻辑

	setCtx(ctx *gin.Context)
	initLog()
}

type Flow struct {
	ctx      *gin.Context
	httpCode int
	logger   *logx.Gin
}

func (f *Flow) setCtx(ctx *gin.Context) {
	f.ctx = ctx
	f.ctx.Set(authorCtxKey, definex.Author)
	f.ctx.Set(definex.TraceId, fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", time.Now().UnixNano()/1e6)))))
}

func (f *Flow) initLog() {
	if f.logger == nil {
		f.logger = logx.NewGin()
	}
}

func (f *Flow) GetAuthor() string {
	return f.ctx.GetString(authorCtxKey)
}

func (f *Flow) Handle() {
	panic("implement func Handle")
}

func (f *Flow) Validate() {
	panic("implement func Validate")
}

func (f *Flow) Logic() {
	panic("implement func Logic")
}

func (f *Flow) ReturnJson(data *replyx.T) {
	var code int
	if f.httpCode == 0 {
		code = http.StatusOK
	} else {
		code = f.httpCode
	}
	f.ctx.JSON(code, data)
}

func (f *Flow) SetHttpCode(code int) {
	f.httpCode = code
}

func (f *Flow) PrintErrorf(format string, a ...interface{}) {
	logf(ErrorLevel, f.ctx, format, a...)
}

func (f *Flow) PrintInfof(format string, a ...interface{}) {
	logf(InfoLevel, f.ctx, format, a...)
}
