package other

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetProjectName 获取项目名称
// 原理: 获取go.mod第一行module后面的内容, 例如: module github.com/liuxiaobopro/gobox
// 例如: 参考 go_test.go ==> TestGetProjectName()
func GetProjectName(path string) (string, error) {
	// 打开文件
	file, err := os.Open(path + "/go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 使用bufio读取文件的第一行
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		if len(firstLine) > 7 {
			return strings.Trim(firstLine[7:], " "), nil
		} else {
			return "", fmt.Errorf("go.mod文件格式错误")
		}
	}
	return "", nil
}
