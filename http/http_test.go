package http

import (
	"testing"
)

func TestClient_Get(t *testing.T) {
	// client := Client{
	// 	Url: "http://localhost:8080/goodsInfo",
	// }

	// resp, err := client.Get()
	// if err != nil {
	// 	t.Error(err)
	// }

	// t.Log(string(resp))
}

func TestClient_Post(t *testing.T) {
}

func TestIsValidUrl(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				str: "http://www.baidu.com",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				str: "https://www.baidu.com",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				str: "www.baidu.com",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				str: "1dd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUrl(tt.args.str); got != tt.want {
				t.Errorf("IsValidUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
