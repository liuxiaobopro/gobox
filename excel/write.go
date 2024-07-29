package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type WriteOption func(*Write)

func WithHead(head []string) WriteOption {
	return func(w *Write) {
		w.Head = head
	}
}

func WithHeadWidth(hw []int) WriteOption {
	return func(w *Write) {
		w.HeadWidth = hw
	}
}

func WithData(data [][]string) WriteOption {
	return func(w *Write) {
		w.Data = data
	}
}

func WithFilepath(filepath string) WriteOption {
	return func(w *Write) {
		w.Filepath = filepath
	}
}

func WithWriteSheet(sheet string) WriteOption {
	return func(w *Write) {
		w.Sheet = sheet
	}
}

func NewWrite(opts ...WriteOption) *Write {
	w := &Write{}
	for _, opt := range opts {
		opt(w)
	}

	if w.Sheet == "" {
		w.Sheet = "Sheet1"
	}

	return w
}

func (w *Write) Write() (string, error) {
	// 创建文件
	f := excelize.NewFile()

	// 创建工作表
	if _, err := f.NewSheet(w.Sheet); err != nil {
		return "", err
	}

	// 设置表头
	for i, v := range w.Head {
		col := 'A' + i
		// accii to string
		_ = f.SetCellValue(w.Sheet, fmt.Sprintf("%c1", col), v)
		// 设置宽度
		if len(w.HeadWidth) > 0 {
			_ = f.SetColWidth(w.Sheet, fmt.Sprintf("%c", col), fmt.Sprintf("%c", col), float64(w.HeadWidth[i]))
		}
	}

	// 设置数据
	for i, v := range w.Data {
		for j, k := range v {
			col := 'A' + j
			_ = f.SetCellValue(w.Sheet, fmt.Sprintf("%c%d", col, i+2), k)
		}
	}

	// 保存文件
	if err := f.SaveAs(w.Filepath); err != nil {
		return "", err
	}

	return w.Filepath, nil
}
