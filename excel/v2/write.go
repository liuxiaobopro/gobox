package excel

import (
	"fmt"
	"strings"

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
		var cell string

		if i < 26 {
			cell = fmt.Sprintf("%c1", 'A'+i)
		} else if i < 52 {
			cell = fmt.Sprintf("A%c1", 'A'+i-26)
		} else if i < 78 {
			cell = fmt.Sprintf("B%c1", 'A'+i-52)
		} else {
			w.err = fmt.Errorf("太多列啦")
			return w
		}

		// accii to string
		_ = w.f.SetCellValue(s.Name, cell, v)
		// 设置宽度
		if len(s.HeadWidth) > 0 {
			col := strings.TrimRight(cell, "1")
			if err := w.f.SetColWidth(s.Name, col, col, float64(s.HeadWidth[i])); err != nil {
				w.err = err
				return w
			}
		}
	}

	// 设置数据
	for i, v := range s.Data {
		for j, k := range v {
			var cell string

			if j < 26 {
				cell = fmt.Sprintf("%c%d", 'A'+j, i+2)
			} else if j < 52 {
				cell = fmt.Sprintf("A%c%d", 'A'+j-26, i+2)
			} else if j < 78 {
				cell = fmt.Sprintf("B%c%d", 'A'+j-52, i+2)
			} else {
				w.err = fmt.Errorf("太多列啦")
				return w
			}

			_ = w.f.SetCellValue(s.Name, cell, k)
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
