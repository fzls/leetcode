package leetcode

import "testing"

// 2019/10/08 0:17 by fzls

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 0}}, want: 3},
		{name: "test", args: struct{ nums []int }{nums: []int{3, 4, -1, 1}}, want: 2},
		{name: "test", args: struct{ nums []int }{nums: []int{3, 4, -1, 1, 2}}, want: 5},
		{name: "test", args: struct{ nums []int }{nums: []int{7, 8, 9, 11, 12}}, want: 1},
		{name: "test", args: struct{ nums []int }{nums: []int{0, 2, 2, 1, 1}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
