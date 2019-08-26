package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		struct {
			name string
			args args
			want int
		}{name: "1", args: struct{ x int }{x: 321}, want: 123},
		struct {
			name string
			args args
			want int
		}{name: "2", args: struct{ x int }{x: -321}, want: -123},
		struct {
			name string
			args args
			want int
		}{name: "1", args: struct{ x int }{x: 0}, want: 0},
		struct {
			name string
			args args
			want int
		}{name: "1", args: struct{ x int }{x: 1534236469}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}