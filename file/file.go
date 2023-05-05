package file

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

// Append 往文件内追加内容
func Append(path string, content string) error {
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

// Has 判断文件内容是否包含某个字符串
func Has(path string, content string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), content) {
			return true, nil
		}
	}

	return false, nil
}

// FindStringsBetween 从字符串中查找两个字符串之间的内容
func FindStringsBetween(str, reg string) string {
	re := regexp.MustCompile(reg)
	matches := re.FindStringSubmatch(str)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
