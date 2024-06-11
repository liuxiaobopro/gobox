package log

import (
	"bytes"
	"fmt"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	definex "github.com/liuxiaobopro/gobox/define"
)

type Mode string

const (
	DefaultMode Mode = "default"
	FileMode    Mode = "file"
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

type Gin struct {
	Mode              Mode
	Level             Level
	Sign              string
	IsCloseColor      bool
	FilePathCutLength int // 截掉文件路径头部长度
}

type GinOption func(c *Gin)

func WithMode(mode Mode) GinOption {
	return func(c *Gin) {
		c.Mode = mode
	}
}

func WithLevel(level Level) GinOption {
	return func(c *Gin) {
		c.Level = level
	}
}

func WithSign(sign string) GinOption {
	return func(c *Gin) {
		c.Sign = sign
	}
}

func WithIsCloseColor(closeColor bool) GinOption {
	return func(c *Gin) {
		c.IsCloseColor = closeColor
	}
}

func WithFilePathCutLength(length int) GinOption {
	return func(c *Gin) {
		c.FilePathCutLength = length
	}
}

func NewGin(op ...GinOption) *Gin {
	var c = &Gin{
		IsCloseColor: false,
	}
	for _, o := range op {
		o(c)
	}

	if c.Mode == "" {
		c.Mode = DefaultMode
	}

	if c.Level == "" {
		c.Level = InfoLevel
	}

	if c.Sign == "" {
		c.Sign = definex.Project
	}

	return c
}

func (conf *Gin) Debugf(c *gin.Context, format string, a ...interface{}) {
	conf.logf(DebugLevel, c, format, a...)
}

func (conf *Gin) Infof(c *gin.Context, format string, a ...interface{}) {
	conf.logf(InfoLevel, c, format, a...)
}

func (conf *Gin) Warnf(c *gin.Context, format string, a ...interface{}) {
	conf.logf(WarnLevel, c, format, a...)
}

func (conf *Gin) Errorf(c *gin.Context, format string, a ...interface{}) {
	conf.logf(ErrorLevel, c, format, a...)
}

func (conf *Gin) Panicf(c *gin.Context, format string, a ...interface{}) {
	conf.logf(PanicLevel, c, format, a...)
}

func (conf *Gin) Fatalf(c *gin.Context, format string, a ...interface{}) {
	conf.logf(FatalLevel, c, format, a...)
}

func (conf *Gin) SetCloseColor(closeColor bool) {
	conf.IsCloseColor = closeColor
}

func (conf *Gin) logf(level Level, c *gin.Context, format string, a ...interface{}) {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "[%s] | %s | %s ", conf.Sign, level, time.Now().Format("2006-01-02 15:04:05"))

	_, file, line, ok := runtime.Caller(2)
	file = file[conf.FilePathCutLength:]
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

	if !conf.IsCloseColor {
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
	}

	fmt.Print(buf.String())
	fmt.Print("\033[0m") // 还原颜色
	fmt.Print("\033[0m") // 还原颜色
	fmt.Print("\033[0m") // 还原颜色
}
