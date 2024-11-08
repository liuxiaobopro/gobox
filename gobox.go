package gobox

import (
	"reflect"

	"github.com/spf13/cast"
)

// Select 模拟三元运算符
func Select[T any](exp bool, yes, no T) T {
	if exp {
		return yes
	}
	return no
}

// SetDefault 给结构体设置默认值
//
// 用法:
//
//	type person struct {
//		Name   string      `default:"test"`
//		Age    int         `default:"18"`
//		Status enum.String `default:"1111"`
//	}
func SetDefault(in interface{}) {
	if in == nil {
		return
	}
	v := reflect.ValueOf(in).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag

		tagValue := tag.Get("default")
		if tagValue != "" {
			if !field.IsZero() {
				continue
			}

			switch field.Kind() {
			case reflect.String:
				field.SetString(tagValue)

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				field.SetInt(cast.ToInt64(tagValue))

			case reflect.Float32, reflect.Float64:
				field.SetFloat(cast.ToFloat64(tagValue))

			case reflect.Bool:
				field.SetBool(cast.ToBool(tagValue))

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				field.SetUint(cast.ToUint64(tagValue))
			}

		}
	}
}
