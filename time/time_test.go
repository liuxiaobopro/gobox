package time

import (
	"testing"
)

func TestIntToString(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{t: 1681834378},
			want: "2023-04-19 00:12:58",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToString(tt.args.t); got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{t: "2023-04-19 00:12:58"},
			want: 1681834378,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt(tt.args.t); got != tt.want {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToStringDate(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{t: 1681834378},
			want: "2023-04-19",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToStringDate(tt.args.t); got != tt.want {
				t.Errorf("IntToStringDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToStringTime(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{t: 1681834378},
			want: "00:12:58",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToStringTime(tt.args.t); got != tt.want {
				t.Errorf("IntToStringTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
