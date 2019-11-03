package leetcode

import "testing"

// 2019/11/03 20:31 by fzls

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct {
			m int
			n int
		}{m: 3, n: 2}, want: 3},
		{name: "test", args: struct {
			m int
			n int
		}{m: 7, n: 3}, want: 28},
		{name: "test", args: struct {
			m int
			n int
		}{m: 1, n: 1}, want: 1},
		{name: "test", args: struct {
			m int
			n int
		}{m: 1, n: 111}, want: 1},
		{name: "test", args: struct {
			m int
			n int
		}{m: 111, n: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
