package log

import "fmt"

func Infof(format string, a ...any) {
	_, _ = fmt.Printf(format, a...)
}

func Errorf(format string, a ...any) {
	_ = fmt.Errorf(format, a...)
}
