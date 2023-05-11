package string

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/liuxiaobopro/gobox/crypto"
)

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

// FirstUp 首字母大写
func FirstUp(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}

// ReplaceCharAfterSpecifiedCharUp 替换指定字符后面的字符为大驼峰
// 例如：ReplaceCharAfterSpecifiedCharUp("abc/def/ghi", '/') => "AbcDefGhi"
func ReplaceCharAfterSpecifiedCharUp(s, c string) (out string) {
	arr := strings.Split(s, c)
	for _, v := range arr {
		v = strings.ToLower(v)
		out += FirstUp(v)
	}
	return
}

// ReplaceCharAfterSpecifiedCharLow 替换指定字符后面的字符为小驼峰
// 例如：ReplaceCharAfterSpecifiedCharLow("abc/def/ghi", '/') => "abcDefGhi"
func ReplaceCharAfterSpecifiedCharLow(s, c string) (out string) {
	arr := strings.Split(s, c)
	for k, v := range arr {
		if k == 0 {
			out += v
			continue
		}
		v = strings.ToLower(v)
		out += FirstUp(v)
	}
	return
}

// Rand 生成随机字符串
func Rand(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

type SafeRand struct {
	Str  string
	lock sync.Mutex
}

// Rand 生成随机字符串
func (sr *SafeRand) Rand() string {
	sr.lock.Lock()
	defer sr.lock.Unlock()
	return crypto.Md5(strconv.Itoa(int(time.Now().UnixNano())) + sr.Str)
}
