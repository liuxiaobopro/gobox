package number

import (
	"math/rand"
	"time"
)

// RandomInt 获取随机数(左开右闭)
func RandomInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Intn(max-min)
}
