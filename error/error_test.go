package error

import (
	"testing"
)

func TestNew(t *testing.T) {
	e := New("test11111")

	t.Log(e)
	t.Log(e.Error())
	t.Log(e.(*T).Msg.String())
	t.Log(e.(*T))
	t.Log(e.(*T).Value())
}
