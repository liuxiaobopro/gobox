package file

import "testing"

func TestAppendFile(t *testing.T) {
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
			if err := AppendFile(tt.args.path, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("AppendFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
