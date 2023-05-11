package log

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/gobox/define"
	"github.com/liuxiaobopro/gobox/time"
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
	FatalLevel Level = "fatal"
	PanicLevel Level = "panic"
)

type Gin struct {
	Mode  Mode
	Level Level
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

func NewGin(op ...GinOption) *Gin {
	var c = &Gin{}
	for _, o := range op {
		o(c)
	}

	if c.Mode == "" {
		c.Mode = DefaultMode
	}

	if c.Level == "" {
		c.Level = InfoLevel
	}

	return c
}

func (conf *Gin) Infof(c *gin.Context, format string, a ...any) {
	str := fmt.Sprintf("[Gobox] | %s | %s ", InfoLevel, time.NowTimeStr())
	_, file, line, ok := runtime.Caller(1)
	if ok {
		str += fmt.Sprintf("| %s:%d ", file, line)
	}
	if c != nil {
		if v, has := c.Get(define.TraceId); has {
			str += fmt.Sprintf("| %s:%s ", define.TraceId, v.(string))
		}
	}
	fmt.Printf(fmt.Sprintf("%s | %s %s", str, format, "\n"), a...)
}

func (conf *Gin) Errorf(c *gin.Context, format string, a ...any) {
	str := fmt.Sprintf("[Gobox] | %s | %s ", ErrorLevel, time.NowTimeStr())
	_, file, line, ok := runtime.Caller(1)
	if ok {
		str += fmt.Sprintf("| %s:%d ", file, line)
	}
	if c != nil {
		if v, has := c.Get(define.TraceId); has {
			str += fmt.Sprintf("| %s:%s ", define.TraceId, v.(string))
		}
	}
	fmt.Printf(fmt.Sprintf("%s | %s %s", str, format, "\n"), a...)
}
