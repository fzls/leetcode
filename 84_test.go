package leetcode

import "testing"

// 2019/11/21 22:24 by fzls

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ heights []int }{heights: []int{2, 1, 5, 6, 2, 3}}, want: 10},
		{name: "test", args: struct{ heights []int }{heights: []int{}}, want: 0},
		{name: "test", args: struct{ heights []int }{heights: []int{2}}, want: 2},
		{name: "test", args: struct{ heights []int }{heights: []int{2, 2, 2, 2}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea(%v) = %v, want %v", tt.args.heights, got, tt.want)
			}
		})
	}
}
