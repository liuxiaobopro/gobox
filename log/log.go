package log

import (
	"context"
	"fmt"
	"runtime"
)

func Infof(c context.Context, format string, a ...any) {
	str := "[Gobox-Info]"
	_, file, line, ok := runtime.Caller(1)
	if ok {
		str += fmt.Sprintf(" %s:%d ", file, line)
	}
	s := str + format
	fmt.Printf(s, a...)
}

func Errorf(c context.Context, format string, a ...any) {
	str := "[Gobox-Error]"
	_, file, line, ok := runtime.Caller(1)
	if ok {
		str += fmt.Sprintf(" %s:%d ", file, line)
	}
	s := str + format
	fmt.Printf(s, a...)
}
