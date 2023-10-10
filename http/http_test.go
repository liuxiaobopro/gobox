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

func TestIsValidUrl1(t *testing.T) {
	url := "https://qiniu.xxx.com/VRvideo/HANGKONGMUJIAN/%E7%BE%8E%E5%86%9B%E7%8E%B0%E5%BD%B9%E6%9C%80%E5%85%88%E8%BF%9B%E7%9A%84%E5%B0%BC%E7%B1%B3%E5%85%B9%E7%BA%A7%E6%A0%B8%E5%8A%A8%E5%8A%9B%E8%88%AA%E7%A9%BA%E6%AF%8D%E8%88%B0%E7%BD%97%E7%BA%B3%E5%BE%B7%E9%87%8C%E6%A0%B9%E5%8F%B7%E8%88%AA%E6%AF%8D%E6%88%98%E6%96%97%E5%8A%9B%E6%80%8E%E4%B9%88%E6%A0%B7%EF%BC%9F.mp4"

	t.Log(IsValidUrl(url))
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
