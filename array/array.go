package array

import (
	typesx "github.com/liuxiaobopro/golib/types"
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
