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
	SetHttpCode(code int)                           // 设置 http code
	ReturnJson(data *replyx.T)                      // 返回 json
	ReturnSucc(obj interface{})                     // 返回成功
	PrintErrorf(format string, args ...interface{}) // 打印错误日志
	PrintInfof(format string, args ...interface{})  // 打印信息日志
	ShouldBind(obj interface{}) error               // 绑定form
	ShouldBindJSON(obj interface{}) error           // 绑定json
	GetCtx() *gin.Context                           // 获取 gin 上下文
	SetReq(obj interface{})                         // 绑定请求参数
	GetReq() interface{}                            // 获取请求参数

	FlowHandle()   // 业务逻辑-控制器句柄
	FlowValidate() // 业务逻辑-参数校验
	FlowLogic()    // 业务逻辑-业务逻辑

	setCtx(ctx *gin.Context)
	initLog()
}

type Flow struct {
	ctx      *gin.Context
	httpCode int
	logger   *logx.Gin
	req      interface{}
}

func (f *Flow) initLog() {
	if f.logger == nil {
		f.logger = logx.NewGin()
	}
}

func (f *Flow) GetAuthor() string {
	return f.ctx.GetString(authorCtxKey)
}

func (f *Flow) FlowHandle() {
	panic("implement func Handle")
}

func (f *Flow) FlowValidate() {
	panic("implement func Validate")
}

func (f *Flow) FlowLogic() {
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

func (f *Flow) ReturnSucc(data interface{}) {
	f.ReturnJson(replyx.Succ(data))
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

func (f *Flow) setCtx(ctx *gin.Context) {
	f.ctx = ctx
	f.ctx.Set(authorCtxKey, definex.Author)
	f.ctx.Set(definex.TraceId, fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", time.Now().UnixNano()/1e6)))))
}

func (f *Flow) GetCtx() *gin.Context {
	return f.ctx
}

func (f *Flow) ShouldBind(obj interface{}) error {
	return f.ctx.ShouldBind(obj)
}

func (f *Flow) SetReq(obj interface{}) {
	f.req = obj
}

func (f *Flow) GetReq() interface{} {
	return f.req
}

func (f *Flow) ShouldBindJSON(obj interface{}) error {
	return f.ctx.ShouldBindJSON(obj)
}
