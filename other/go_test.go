package other

import (
	"testing"
)

func TestGetProjectName(t *testing.T) {
	path := "D:\\1_lxb2\\1_project\\gobox"

	// // 获取可执行文件的路径
	// exePath, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }

	// // 获取可执行文件所在的目录
	// exeDir := filepath.Dir(exePath)

	// // 获取项目目录
	// projectDir := filepath.Dir(exeDir)

	// t.Logf("projectDir: [%s]", projectDir)
	// path := projectDir

	data, err := GetProjectName(path)
	if err != nil {
		t.Error(err)
	}

	t.Logf("module: [%s]", data)
}
