package types

import "time"

// All 支持的所有类型
type All interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool | []byte | time.Time
}

// All1 支持的所有类型(不包含[]byte)
type All1 interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool | time.Time
}

// IntUintFloat 支持的所有整数和浮点数类型
type IntUintFloat interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

