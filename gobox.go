package gobox

// Select 模拟三元运算符
func Select[T any](exp bool, yes, no T) T {
	if exp {
		return yes
	}
	return no
}
