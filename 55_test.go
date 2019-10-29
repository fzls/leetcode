package leetcode

import "testing"

// 2019/10/30 0:42 by fzls

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{2, 3, 1, 1, 4}}, want: true},
		{name: "test", args: struct{ nums []int }{nums: []int{3, 2, 1, 0, 4}}, want: false},
		{name: "test", args: struct{ nums []int }{nums: []int{0, 1}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
