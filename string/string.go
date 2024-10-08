package string

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/liuxiaobopro/gobox/crypto"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	if len(s) == 1 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}

// FirstLow 首字母小写
func FirstLow(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}
	return strings.ToLower(s[0:1]) + s[1:]
}

// ReplaceCharAfterSpecifiedCharUp 替换指定字符后面的字符为大驼峰
// 例如：ReplaceCharAfterSpecifiedCharUp("abc/def/ghi", '/') => "AbcDefGhi"
func ReplaceCharAfterSpecifiedCharUp(s, c string) (out string) {
	arr := strings.Split(s, c)
	for _, v := range arr {
		v1 := v[0:1]
		v2 := v[1:]
		v = strings.ToLower(v1)
		out += FirstUp(v + v2)
	}
	return
}

// CamelCase 转大驼峰
func CamelCase(s string) string {
	str := LowerCamelCase(s)

	return FirstUp(str)
}

// LowerCamelCase 转小驼峰
func LowerCamelCase(s string) string {
	words := regexp.MustCompile(`[-_\s]+`).Split(s, -1)
	lowerCamelCase := ""
	for i, word := range words {
		if i == 0 {
			lowerCamelCase += strings.ToLower(word)
		} else {
			c := cases.Title(language.Und, cases.NoLower).String(word)
			lowerCamelCase += c
		}
	}
	return lowerCamelCase
}

// ReplaceCharAfterSpecifiedCharLow 替换指定字符后面的字符为小驼峰
// 例如：ReplaceCharAfterSpecifiedCharLow("abc/def/ghi", '/') => "abcDefGhi"
func ReplaceCharAfterSpecifiedCharLow(s, c string) (out string) {
	// arr := strings.Split(s, c)
	// for k, v := range arr {
	// 	if k == 0 {
	// 		out += v
	// 		continue
	// 	}
	// 	v = strings.ToLower(v)
	// 	out += FirstUp(v)
	// }
	str := ReplaceCharAfterSpecifiedCharUp(s, c)
	out = FirstLow(str)
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

// RandV1 生成随机字符串
func RandFor(l int, seed int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandStrArr 生成随机字符串数组
// @param l 长度
// @param n 数量
func RandStrArr(l int, n int) []string {
	var out []string
	for i := 0; i < n; i++ {
		out = append(out, RandFor(l, time.Now().UnixNano()+int64(i)))
	}
	return out
}

// RandInt 生成随机数字字符串
func RandInt(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成随机小写字母字符串
func RandLower(l int) string {
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成随机大写字母字符串
func RandUpper(l int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成自定义字符串
func RandCustom(l int, str string) string {
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

// IsNumber 判断是否是数字或者小数
func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// ToAccii 将大写字母转换为accii码
func ToAccii(s string) int {
	return int([]rune(s)[0])
}

// StrToAccii 将accii码转换为大写字母
func AcciiToStr(s int) string {
	return string(rune(s))
}

// UniqueFileName 生成唯一文件名
func UniqueFileName(f *multipart.FileHeader) string {
	fileSuffix := path.Ext(f.Filename)
	fileName := strings.TrimSuffix(f.Filename, fileSuffix)
	fileName = fmt.Sprintf("%s_%s%s", fileName, crypto.Md5(strconv.Itoa(int(time.Now().UnixNano()))), fileSuffix)
	return fileName
}

// HasChinese 判断字符串是否存在中文
func HasChinese(s string) bool {
	for _, v := range s {
		if v > 0x4e00 && v < 0x9fff {
			return true
		}
	}
	return false
}

// StringValueIsEqual 比较两个字符串是否相等(string 转 float64 作比较)
func StringValueIsEqual(a, b string) bool {
	// 将字符串转换为 float64
	num1, err1 := strconv.ParseFloat(a, 64)
	num2, err2 := strconv.ParseFloat(b, 64)

	if err1 != nil || err2 != nil {
		// 解析失败，不相等
		return false
	}

	// 使用容差值比较两个浮点数
	const tolerance = 1e-9
	return (num1-num2) < tolerance && (num2-num1) < tolerance
}
