package file

import (
	"testing"
)

func TestAppend(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{"./test.txt", "test\n"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Append(tt.args.path, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHas(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"test1", args{"./test.txt", "test"}, true, false},
		{"test2", args{"./test.txt", "test1"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Has(tt.args.path, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Has() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindStringsBetween(t *testing.T) {
	str := `
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
    `
	reg := "import \\((?s)(.*?)\\)"
	got := FindStringsBetween(str, reg)
	t.Log("got: ", got)
}
