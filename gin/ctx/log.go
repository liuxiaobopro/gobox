package ctx

import (
	"bytes"
	"fmt"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	definex "github.com/liuxiaobopro/gobox/define"
)

type Level string

const (
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	WarnLevel  Level = "warn"
	ErrorLevel Level = "error"
	PanicLevel Level = "panic"
	FatalLevel Level = "fatal"
)

func logf(level Level, c *gin.Context, format string, a ...interface{}) {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "[Flow] | %s | %s ", level, time.Now().Format("2006-01-02 15:04:05"))

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
