package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Write struct {
	Filepath string // 文件路径

	f           *excelize.File // 文件
	sheets      []*Sheet       // 工作表
	isDelSheet1 bool           // 是否删除默认工作表

	err error
}

type writeOption func(*Write)

func NewExcel(filepath string, opts ...writeOption) *Write {
	w := &Write{
		Filepath: filepath,

		f: excelize.NewFile(),
	}

	for _, opt := range opts {
		opt(w)
	}

	return w
}

func (w *Write) DelSheet1() *Write {
	w.isDelSheet1 = true

	return w
}

type Sheet struct {
	Name      string     // 工作表名称
	Head      []string   // 表头
	HeadWidth []int      // 表头宽度
	Data      [][]string // 数据
}

type sheetOption func(*Sheet)

func WithSheetHead(head []string) sheetOption {
	return func(s *Sheet) {
		s.Head = head
	}
}

func WithSheetHeadWidth(hw []int) sheetOption {
	return func(s *Sheet) {
		s.HeadWidth = hw
	}
}

func WithSheetData(data [][]string) sheetOption {
	return func(s *Sheet) {
		s.Data = data
	}
}

func NewSheet(name string, opts ...sheetOption) *Sheet {
	s := &Sheet{
		Name: name,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (w *Write) AddSheet(s *Sheet) *Write {
	w.sheets = append(w.sheets, s)

	// 创建工作表
	if _, err := w.f.NewSheet(s.Name); err != nil {
		w.err = err

		return w
	}

	// 设置表头
	for i, v := range s.Head {
		col := 'A' + i
		// accii to string
		_ = w.f.SetCellValue(s.Name, fmt.Sprintf("%c1", col), v)
		// 设置宽度
		if len(s.HeadWidth) > 0 {
			if err := w.f.SetColWidth(s.Name, fmt.Sprintf("%c", col), fmt.Sprintf("%c", col), float64(s.HeadWidth[i])); err != nil {
				w.err = err
				return w
			}
		}
	}

	// 设置数据
	for i, v := range s.Data {
		for j, k := range v {
			col := 'A' + j
			if err := w.f.SetCellValue(s.Name, fmt.Sprintf("%c%d", col, i+2), k); err != nil {
				w.err = err
				return w
			}
		}
	}

	return w
}

func (w *Write) Save() error {
	if w.err != nil {
		return w.err
	}

	if w.isDelSheet1 {
		_ = w.f.DeleteSheet("sheet1")
		_ = w.f.DeleteSheet("Sheet1")
	}

	// 保存文件
	if err := w.f.SaveAs(w.Filepath); err != nil {
		return err
	}

	return nil
}
