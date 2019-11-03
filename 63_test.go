package leetcode

import "testing"

// 2019/11/03 21:04 by fzls

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}}, want: 2},
		{name: "test", args: struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{
			{1, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}}, want: 0},
		{name: "test", args: struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 1},
		}}, want: 0},
		{name: "test", args: struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{
			{0, 0, 0},
			{0, 0, 1},
			{0, 1, 0},
		}}, want: 0},
		{name: "test", args: struct{ obstacleGrid [][]int }{obstacleGrid: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 1, 0},
		}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
