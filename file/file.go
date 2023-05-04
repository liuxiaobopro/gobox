package file

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ReplaceInDir 替换目录下所有文件中的字符串
func ReplaceInDir(dirPath, oldStr, newStr string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		newContent := strings.Replace(string(content), oldStr, newStr, -1)

		if newContent != string(content) {
			err = ioutil.WriteFile(path, []byte(newContent), info.Mode())
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// AppendFile 往文件内追加内容
func AppendFile(path string, content string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
