package number

import (
	"math/rand"
	"time"
)

// RandomInt 获取随机数(左开右闭)
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
