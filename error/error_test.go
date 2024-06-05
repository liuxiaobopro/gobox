package error

import (
	"testing"

	"github.com/liuxiaobopro/gobox/reply"
)

func TestNew(t *testing.T) {
	e := New("致命错误", WithCode(11), WithData("test22222"))

	t.Log(e)
	t.Log(e.(*reply.T).Code)
	t.Log(e.(*reply.T).Data)
}
