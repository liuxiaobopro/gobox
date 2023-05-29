package excel

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