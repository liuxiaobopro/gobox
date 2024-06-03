package gobox

import (
	"fmt"
	"testing"
)

type person struct {
	Name string
}

func (p person) String() string {
	return fmt.Sprintf("Name: %s", p.Name)
}

func TestSelect(t *testing.T) {
	a := 1
	b := 2

	c := person{Name: "c"}
	d := person{Name: "d"}

	res := Select(a < b, c, d)

	t.Log(res)

	func1 := func() {
		t.Log("func1")
	}

	func2 := func() {
		t.Log("func2")
	}

	Select(a < b, func1, func2)()
}
