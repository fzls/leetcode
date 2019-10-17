package leetcode

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
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
		}{name: "test", args: struct{ nums []int }{nums: []int{2, 3, 1, 1, 4, 1, 2, 3, 1, 1, 2, 5, 7, 1, 2, 3, 5, 8, 1, 2, 3, 4, 1, 8, 2, 1}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
