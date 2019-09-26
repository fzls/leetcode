package leetcode

import "testing"

// 2019/09/27 0:54 by fzls

func Test_countDigitOne(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ n int }{n: 1}, want: 1},
		{name: "test", args: struct{ n int }{n: 5}, want: 1},
		{name: "test", args: struct{ n int }{n: 10}, want: 2},
		{name: "test", args: struct{ n int }{n: 13}, want: 6},
		{name: "test", args: struct{ n int }{n: 55}, want: 16},
		{name: "test", args: struct{ n int }{n: 99}, want: 20},
		{name: "test", args: struct{ n int }{n: 10000}, want: 4001},
		{name: "test", args: struct{ n int }{n: 21345}, want: 18821},
		{name: "test", args: struct{ n int }{n: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigitOne(tt.args.n); got != tt.want {
				t.Errorf("countDigitOne(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
