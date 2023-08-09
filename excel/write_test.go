package excel

import "testing"

func TestWrite_Write(t *testing.T) {
	w := NewWrite(
		WithFilepath("test.xlsx"),
		WithWriteSheet("Sheet1"),
		WithHead([]string{"姓名", "年龄", "性别"}),
		WithData([][]string{
			{"张三", "18", "男"},
			{"李四", "19", "女"},
			{"王五", "20", "男"},
		}),
	)

	if _, err := w.Write(); err != nil {
		t.Error(err)
		return
	}

	t.Log("success")
}
