package array

import (
	"testing"
	"time"
)

func TestRandomSelection(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	randomCount := 4

	t.Log("随机选择的值:", RandomSelection(intSlice, randomCount))
	time.Sleep(1 * time.Second)
	t.Log("随机选择的值:", RandomSelection(intSlice, randomCount))
	time.Sleep(1 * time.Second)
	t.Log("随机选择的值:", RandomSelection(intSlice, randomCount))
	time.Sleep(1 * time.Second)
	t.Log("随机选择的值:", RandomSelection(intSlice, randomCount))
	time.Sleep(1 * time.Second)
	t.Log("随机选择的值:", RandomSelection(intSlice, randomCount))
	time.Sleep(1 * time.Second)
}
