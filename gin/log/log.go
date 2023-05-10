package log

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/gobox/define"
)

func Infof(c *gin.Context, format string, a ...any) {
	str := "[Gobox-Info]"
	_, file, line, ok := runtime.Caller(1)
	if ok {
		str += fmt.Sprintf(" %s:%d ", file, line)
	}
	if c != nil {
		if v, has := c.Get(define.TRACE_ID); has {
			str += fmt.Sprintf(" %s:%s ", define.TRACE_ID, v.(string))
		}
	}
	fmt.Printf(str+format, a...)
}

func Errorf(c *gin.Context, format string, a ...any) {
	str := "[Gobox-Error]"
	_, file, line, ok := runtime.Caller(1)
	if ok {
		str += fmt.Sprintf(" %s:%d ", file, line)
	}
	if c != nil {
		if v, has := c.Get(define.TRACE_ID); has {
			str += fmt.Sprintf(" %s:%s ", define.TRACE_ID, v.(string))
		}
	}
	fmt.Printf(str+format, a...)
}
