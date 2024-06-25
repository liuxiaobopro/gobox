package reply

import (
	"testing"
)

func TestNew(t *testing.T) {
	e1 := New("err1")
	t.Logf("%v", e1)

	e2 := New("err2", WithData("data"))
	t.Logf("%v", e2)

	e3 := New("err3", WithCode(11), WithData("data"))
	t.Logf("%v", e3)
}
