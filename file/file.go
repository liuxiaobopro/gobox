package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
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

// Upload 上传文件
// @param file 文件
// @param fileHeader 文件头
// @param filepath 文件保存路径
// @return 文件路径，文件名，错误
func Upload(file multipart.File, fileHeader *multipart.FileHeader, filepath string, filename string) (string, string, error) {
	// 生成文件名
	fileExt := path.Ext(fileHeader.Filename)
	if filename == "" {
		filename = fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)
	}

	// 生成文件夹
	uploadPath := filepath
	// 创建文件夹
	folder := uploadPath
	if err := os.MkdirAll(uploadPath, 0744); err != nil {
		return "", "", err
	}
	// 生成文件路径
	fPath := fmt.Sprintf("%s/%s", folder, filename)
	fW, err := os.Create(fPath)
	if err != nil {
		return "", "", err
	}
	defer fW.Close()
	// 复制文件，保存到本地
	_, err = io.Copy(fW, file)
	if err != nil {
		return "", "", err
	}

	return fPath, filename, nil
}
