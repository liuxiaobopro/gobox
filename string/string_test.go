package string

import (
	"testing"
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
