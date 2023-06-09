package log

import (
	"bytes"
	"fmt"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	definex "github.com/liuxiaobopro/gobox/define"
)

func CtxDebugf(c *gin.Context, format string, a ...interface{}) {
	ctxlogf(DebugLevel, c, format, a...)
}

func CtxInfof(c *gin.Context, format string, a ...interface{}) {
	ctxlogf(InfoLevel, c, format, a...)
}

func CtxWarnf(c *gin.Context, format string, a ...interface{}) {
	ctxlogf(WarnLevel, c, format, a...)
}

func CtxErrorf(c *gin.Context, format string, a ...interface{}) {
	ctxlogf(ErrorLevel, c, format, a...)
}

func CtxPanicf(c *gin.Context, format string, a ...interface{}) {
	ctxlogf(PanicLevel, c, format, a...)
}

func CtxFatalf(c *gin.Context, format string, a ...interface{}) {
	ctxlogf(FatalLevel, c, format, a...)
}

func ctxlogf(level Level, c *gin.Context, format string, a ...interface{}) {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "[%s] | %s | %s ", definex.Project, level, time.Now().Format("2006-01-02 15:04:05"))

	_, file, line, ok := runtime.Caller(2)
	if ok {
		fmt.Fprintf(&buf, "| %s:%d ", file, line)
	}

	if c != nil {
		if v, has := c.Get(definex.TraceId); has {
			fmt.Fprintf(&buf, "| %s:%s | ", definex.TraceId, v.(string))
		}
	}

	fmt.Fprintf(&buf, format, a...)
	fmt.Fprint(&buf, "\n")

	if level == DebugLevel {
		fmt.Print("\033[32m") // 绿色
	}

	if level == InfoLevel {
		fmt.Print("\033[36m") // 青色
	}

	if level == WarnLevel {
		fmt.Print("\033[33m") // 黄色
	}

	if level == ErrorLevel || level == PanicLevel || level == FatalLevel {
		fmt.Print("\033[31m") // 红色
	}

	fmt.Print(buf.String())
	fmt.Print("\033[0m") // 还原颜色
	fmt.Print("\033[0m") // 还原颜色
	fmt.Print("\033[0m") // 还原颜色
}
