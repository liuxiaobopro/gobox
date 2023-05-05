package any

import (
	"github.com/mitchellh/mapstructure"
)

// AnyToStruct 将任意类型转换为结构体(原仓库: https://github.com/mitchellh/mapstructure)
func AnyToStruct(input interface{}, output interface{}) error {
	return mapstructure.Decode(input, output)
}
