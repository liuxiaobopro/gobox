package log

import (
	"github.com/zeromicro/go-zero/core/logx"
)

func Infof(format string, a ...any) {
	logx.Infof(format, a...)
}

func Errorf(format string, a ...any) {
	logx.Errorf(format, a...)
}
