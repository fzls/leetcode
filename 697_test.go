package leetcode

import "testing"

// 2019/09/26 2:01 by fzls

func Test_findShortestSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 2, 3, 1}}, want: 2},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 2, 3, 1, 4, 2}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
