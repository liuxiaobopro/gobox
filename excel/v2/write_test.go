package excel

import (
	"testing"
)

func TestNewWrite(t *testing.T) {
	excel := NewExcel("./test.xlsx")
	excel.DelSheet1()
	excel.NewSheet("sheet1", []string{"a", "b"}, [][]string{{"1", "2"}})
	excel.NewSheet("sheet21", []string{"c", "d"}, [][]string{{"3", "4"}})
	excel.NewSheet("sheet31", []string{"e", "f"}, [][]string{{"5", "6"}})

	if err := excel.Save(); err != nil {
		t.Fatal(err)
	}

	t.Log("success")
}

func TestEQ(t *testing.T) {
	a := "sheet1"
	b := "Sheet1"

	t.Logf("a==b: %v", a == b)
}
