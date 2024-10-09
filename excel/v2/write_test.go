package excel

import (
	"testing"
)

func TestNewWrite(t *testing.T) {
	excel := NewExcel("./test.xlsx")
	excel.DelSheet1()
	excel.AddSheet(NewSheet("人员1",
		WithSheetHead([]string{"name", "age"}),
		WithSheetHeadWidth([]int{10, 10}),
		WithSheetData([][]string{
			{"tom", "18"},
			{"jerry", "20"},
		}),
	)).AddSheet(NewSheet("人员2",
		WithSheetHead([]string{"name", "age"}),
		WithSheetHeadWidth([]int{10, 10}),
		WithSheetData([][]string{
			{"tom", "18"},
			{"jerry", "20"},
		}),
	))

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
