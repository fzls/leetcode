package leetcode

import "testing"

// 2019/10/29 0:01 by fzls

func Test_totalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ n int }{n: 4}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
