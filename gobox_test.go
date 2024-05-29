package gobox

import (
	"testing"
)

type person struct {
	Name string
}

func TestSelect(t *testing.T) {
	a := 1
	b := 2

	c := person{Name: "c"}
	d := person{Name: "d"}

	res := Select(a < b, c, d)

	t.Log(res)
}
