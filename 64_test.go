package leetcode

import "testing"

// 2019/11/03 21:16 by fzls

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ grid [][]int }{grid: [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}}, want: 7},
		{name: "test", args: struct{ grid [][]int }{grid: [][]int{
			{1, 1, 1},
			{1, 5, 1},
			{4, 2, 1},
		}}, want: 5},
		{name: "test", args: struct{ grid [][]int }{grid: [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}}, want: 5},
		{name: "test", args: struct{ grid [][]int }{grid: [][]int{
			{1},
			{1},
			{1},
		}}, want: 3},
		{name: "test", args: struct{ grid [][]int }{grid: [][]int{}}, want: 0},
		{name: "test", args: struct{ grid [][]int }{grid: [][]int{{1, 1, 1, 1}}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
