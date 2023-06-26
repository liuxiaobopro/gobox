package log

import (
	"bytes"
	"fmt"
	"runtime"
	"time"

	definex "github.com/liuxiaobopro/gobox/define"
)

func Debugf(format string, a ...interface{}) {
	logf(DebugLevel, format, a...)
}

func Infof(format string, a ...interface{}) {
	logf(InfoLevel, format, a...)
}

func Warnf(format string, a ...interface{}) {
	logf(WarnLevel, format, a...)
}

func Errorf(format string, a ...interface{}) {
	logf(ErrorLevel, format, a...)
}

func Panicf(format string, a ...interface{}) {
	logf(PanicLevel, format, a...)
}

func Fatalf(format string, a ...interface{}) {
	logf(FatalLevel, format, a...)
}

func logf(level Level, format string, a ...interface{}) {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "[%s] | %s | %s ", definex.Project, level, time.Now().Format("2006-01-02 15:04:05"))

	_, file, line, ok := runtime.Caller(2)
	if ok {
		fmt.Fprintf(&buf, "| %s:%d ", file, line)
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
}
