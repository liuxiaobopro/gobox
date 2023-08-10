package main

import (
	"math/rand"
	"time"
)

// 随机选择 n 个元素并返回新的切片
func RandomSelection[T any](input []T, n int) []T {
	// 创建一个新的随机数生成器
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// 如果 n 大于输入切片的长度，返回整个输入切片
	if n >= len(input) {
		return input
	}

	// 创建一个新切片用于存放随机选择的元素
	output := make([]T, n)

	// 复制输入切片，避免在原切片上进行修改
	tempInput := make([]T, len(input))
	copy(tempInput, input)

	// 随机选择 n 个元素
	for i := 0; i < n; i++ {
		// 生成一个随机索引
		randomIndex := r.Intn(len(tempInput))

		// 将随机选择的元素添加到输出切片中
		output[i] = tempInput[randomIndex]

		// 从临时切片中移除已选择的元素
		tempInput = append(tempInput[:randomIndex], tempInput[randomIndex+1:]...)
	}

	return output
}
