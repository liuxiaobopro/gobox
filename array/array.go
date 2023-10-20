package array

import (
	"reflect"

	typesx "github.com/liuxiaobopro/gobox/types"
)

// Union 并集
func Union[T typesx.All1](a, b []T) []T {
	set := make(map[T]bool)
	var result []T
	for _, v := range a {
		set[v] = true
	}
	for _, v := range b {
		if !set[v] {
			set[v] = true
		}
	}
	for k := range set {
		result = append(result, k)
	}
	return result
}

// Intersection 交集
func Intersection[T typesx.All1](a, b []T) []T {
	set := make(map[T]bool)
	var result []T
	for _, v := range a {
		set[v] = true
	}
	for _, v := range b {
		if set[v] {
			result = append(result, v)
		}
	}
	return result
}

// Difference 差集
func Difference[T typesx.All1](a, b []T) []T {
	set1 := make(map[T]bool)
	var result []T
	for _, v := range a {
		set1[v] = true
	}
	for _, v := range b {
		if !set1[v] {
			result = append(result, v)
		}
	}

	set2 := make(map[T]bool)
	for _, v := range b {
		set2[v] = true
	}
	for _, v := range a {
		if !set2[v] {
			result = append(result, v)
		}
	}
	return result
}

// IsIn 判断元素是否在数组中
func IsIn[T interface{}](a []T, b T) bool {
	for _, v := range a {
		if reflect.DeepEqual(v, b) {
			return true
		}
	}
	return false
}

//	func IsIn[T typesx.All1](a []T, b T) bool {
//		for _, v := range a {
//			if v == b {
//				return true
//			}
//		}
//		return false
//	}

// Delete 删除数组中的元素
func Delete[T typesx.All1](a []T, b ...T) []T {
	for _, v := range b {
		for i, vv := range a {
			if vv == v {
				a = append(a[:i], a[i+1:]...)
			}
		}
	}
	return a
}

// DeleteByIndex 删除数组中的元素
func DeleteByIndex[T typesx.All1](a []T, b int) []T {
	a = append(a[:b], a[b+1:]...)
	return a
}
