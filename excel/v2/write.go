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
	name      string     // 工作表名称
	head      []string   // 表头
	headWidth []int      // 表头宽度
	data      [][]string // 数据
}

type sheetOption func(*Sheet)

func WithSheetHeadWidth(hw []int) func(*Sheet) {
	return func(s *Sheet) {
		s.headWidth = hw
	}
}

func (w *Write) NewSheet(name string, head []string, data [][]string, opts ...sheetOption) *Sheet {
	s := &Sheet{
		name: name,
		head: head,
		data: data,
	}

	for _, opt := range opts {
		opt(s)
	}

	w.addSheet(s)

	return s
}

func (w *Write) addSheet(s *Sheet) *Write {
	w.sheets = append(w.sheets, s)

	// 创建工作表
	if _, err := w.f.NewSheet(s.name); err != nil {
		w.err = err

		return w
	}

	// 设置表头
	for i, v := range s.head {
		col := 'A' + i
		// accii to string
		_ = w.f.SetCellValue(s.name, fmt.Sprintf("%c1", col), v)
		// 设置宽度
		if len(s.headWidth) > 0 {
			if err := w.f.SetColWidth(s.name, fmt.Sprintf("%c", col), fmt.Sprintf("%c", col), float64(s.headWidth[i])); err != nil {
				w.err = err
				return w
			}
		}
	}

	// 设置数据
	for i, v := range s.data {
		for j, k := range v {
			col := 'A' + j
			if err := w.f.SetCellValue(s.name, fmt.Sprintf("%c%d", col, i+2), k); err != nil {
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
