package string

import "strings"

// Has 判断字符串是否存在某个字符
func Has(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

// IsPrefix 判断字符串是否存在某个前缀
func IsPrefix(s string, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

// IsSuffix 判断字符串是否存在某个后缀
func IsSuffix(s string, suffix string) bool {
	if len(s) < len(suffix) {
		return false
	}
	for i := 0; i < len(suffix); i++ {
		if s[len(s)-len(suffix)+i] != suffix[i] {
			return false
		}
	}
	return true
}

// Count 查询字符串中出现某个字符的次数
func Count(s string, c byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++
		}
	}
	return count
}

// CutStartString 截取字符串中最后一个字符之前的字符串
// 例如：CutStartString("abc/def/ghi", '/') => "abc/def/"
func CutStartString(s string, c rune) string {
	i := strings.LastIndex(s, string(c))
	if i == -1 {
		return s
	}
	return s[0 : i+1]
}

// CutEndString 截取字符串中最后一个字符之后的字符串
// 例如：CutEndString("abc/def/ghi", '/') => "ghi"
func CutEndString(s string, c rune) string {
	i := strings.LastIndex(s, string(c))
	if i == -1 {
		return ""
	}
	return s[i+1:]
}
