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
			fmt.Fprintf(&buf, "| %s:%s ", definex.TraceId, v.(string))
		}
	}

	fmt.Fprintf(&buf, format, a...)
	fmt.Fprint(&buf, "\n")

	fmt.Print(buf.String())
}
