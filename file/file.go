package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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
