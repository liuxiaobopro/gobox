package excel

import (
	"fmt"

	stringx "github.com/liuxiaobopro/gobox/string"
	"github.com/xuri/excelize/v2"
)

type Row int // 行
func (r Row) ToInt() int {
	return int(r)
}

type Col string // 列
func (c Col) ToString() string {
	return string(c)
}

type Cell struct {
	Row Row
	Col Col
} // 单元格

type Read struct {
	file  string // 文件
	Lu    Cell   // 左上单元格坐标
	Rd    Cell   // 右下单元格坐标
	Sheet string // 工作表名称
}

type ReadOption func(*Read)

func WithFile(file string) ReadOption {
	return func(r *Read) {
		r.file = file
	}
}

func WithLu(lu Cell) ReadOption {
	return func(r *Read) {
		r.Lu = lu
	}
}

func WithRd(rd Cell) ReadOption {
	return func(r *Read) {
		r.Rd = rd
	}
}

func WithSheet(sheet string) ReadOption {
	return func(r *Read) {
		r.Sheet = sheet
	}
}

func NewRead(opts ...ReadOption) *Read {
	r := &Read{}
	for _, opt := range opts {
		opt(r)
	}

	if r.Lu.Row == 0 || r.Lu.Col == "" || r.Rd.Col == "" {
		panic("左上单元格坐标不能为空")
	}

	if r.Sheet == "" {
		r.Sheet = "Sheet1"
	}

	return r
}

// Read 读取excel文件
func (r *Read) Read() ([][]string, error) {
	var (
		out [][]string
	)

	if r.Rd.Row == 0 {
		row, err := r.GetMaxRowNum()
		if err != nil {
			return nil, err
		}
		r.Rd.Row = Row(row)
	}

	// 读取文件
	f, err := excelize.OpenFile(r.file)
	if err != nil {
		panic(err)
	}

	// 获取工作表中指定单元格的值
	for i := r.Lu.Row.ToInt(); i <= r.Rd.Row.ToInt(); i++ {
		var row []string
		for j := r.Lu.Col.ToString(); j <= r.Rd.Col.ToString(); j = stringx.AcciiToStr(stringx.ToAccii(j) + 1) {
			cell := fmt.Sprintf("%s%d", j, i)
			// 获取单元格内容
			value, err := f.GetCellValue(r.Sheet, cell)
			if err != nil {
				return nil, err
			}
			row = append(row, value)
		}
		out = append(out, row)
	}

	return out, nil
}

// GetMaxRowNum 获取最大行数
func (r *Read) GetMaxRowNum() (int, error) {
	var (
		out int
	)

	// 读取文件
	f, err := excelize.OpenFile(r.file)
	if err != nil {
		panic(err)
	}

	// 获取工作表中指定单元格的值
	for i := r.Lu.Row.ToInt(); i <= r.Rd.Row.ToInt(); i++ {
		cell := fmt.Sprintf("A%d", i)
		// 获取单元格内容
		value, err := f.GetCellValue(r.Sheet, cell)
		if err != nil {
			return 0, err
		}
		if value == "" {
			break
		}
		out++
	}

	return out, nil
}
