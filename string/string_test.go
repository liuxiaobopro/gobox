package string

import (
	"fmt"
	"testing"
	"time"
)

func TestHas(t *testing.T) {
	type args struct {
		s string
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: "abc",
				c: 'a',
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "abc",
				c: '=',
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				s: "a=c",
				c: '=',
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Has(tt.args.s, tt.args.c); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrefix(t *testing.T) {
	type args struct {
		s      string
		prefix string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s:      "abc",
				prefix: "a",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s:      "abc",
				prefix: "b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrefix(tt.args.s, tt.args.prefix); got != tt.want {
				t.Errorf("IsPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSuffix(t *testing.T) {
	type args struct {
		s      string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s:      "abc",
				suffix: "a",
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				s:      "abc",
				suffix: "c",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSuffix(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("IsSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceCharAfterSpecifiedCharUp(t *testing.T) {
	type args struct {
		s string
		c string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: "abc/def/ghi",
				c: "/",
			},
			want: "AbcDefGhi",
		},
		{
			name: "test2",
			args: args{
				s: "abc/def/ghiJKL",
				c: "/",
			},
			want: "AbcDefGhiJKL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceCharAfterSpecifiedCharUp(tt.args.s, tt.args.c); got != tt.want {
				t.Errorf("ReplaceCharAfterSpecifiedCharUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceCharAfterSpecifiedCharLow(t *testing.T) {
	type args struct {
		s string
		c string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "test1",
			args: args{
				s: "abc/def/ghi",
				c: "/",
			},
			wantOut: "abcDefGhi",
		},
		{
			name: "test2",
			args: args{
				s: "app/demo/demo",
				c: "/",
			},
			wantOut: "appDemoDemo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := ReplaceCharAfterSpecifiedCharLow(tt.args.s, tt.args.c); gotOut != tt.wantOut {
				t.Errorf("ReplaceCharAfterSpecifiedCharLow() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestSafeRand_Rand(t *testing.T) {
	sd := &SafeRand{
		Str: Rand(10),
	}

	t.Log(sd.Str + "\n")
	t.Log(sd.Rand() + "\n")
}

func TestIsNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: "123",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "123.3",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "123.3d",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.s); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		code := Rand(10)
		fmt.Println("code: ", code)
	}
}

func TestRandFor(t *testing.T) {
	for i := 0; i < 10; i++ {
		code := RandFor(10, time.Now().UnixNano()+int64(i))
		fmt.Println("code: ", code)
	}
}

func TestRandStrArr(t *testing.T) {
	fmt.Println("RandStrArr: ", RandStrArr(20, 10))
}

func TestHasChinese(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: "abc",
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				s: "abc中文",
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				s: "中文",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasChinese(tt.args.s); got != tt.want {
				t.Errorf("HasChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringValueIsEqual(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				a: "1.1",
				b: "1.1",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				a: "1.1",
				b: "1.2",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				a: "1.1",
				b: "1.10000000000",
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				a: "1.1",
				b: "1.1.1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringValueIsEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("StringValueIsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
