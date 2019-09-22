package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{1}}, want: []int{1}},

		{name: "test", args: struct{ nums []int }{nums: []int{1, 2}}, want: []int{2, 1}},
		{name: "test", args: struct{ nums []int }{nums: []int{2, 1}}, want: []int{1, 2}},

		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 3}}, want: []int{1, 3, 2}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 3, 2}}, want: []int{2, 1, 3}},
		{name: "test", args: struct{ nums []int }{nums: []int{2, 1, 3}}, want: []int{2, 3, 1}},
		{name: "test", args: struct{ nums []int }{nums: []int{2, 3, 1}}, want: []int{3, 1, 2}},
		{name: "test", args: struct{ nums []int }{nums: []int{3, 1, 2}}, want: []int{3, 2, 1}},
		{name: "test", args: struct{ nums []int }{nums: []int{3, 2, 1}}, want: []int{1, 2, 3}},

		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 3, 4}}, want: []int{1, 2, 4, 3}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 4, 3}}, want: []int{1, 3, 2, 4}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 3, 2, 4}}, want: []int{1, 3, 4, 2}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 3, 4, 2}}, want: []int{1, 4, 2, 3}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 4, 2, 3}}, want: []int{1, 4, 3, 2}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 4, 3, 2}}, want: []int{2, 1, 3, 4}},
		{name: "test", args: struct{ nums []int }{nums: []int{2, 1, 3, 4}}, want: []int{2, 1, 4, 3}},
		{name: "test", args: struct{ nums []int }{nums: []int{2, 1, 4, 3}}, want: []int{2, 3, 1, 4}},
		{name: "test", args: struct{ nums []int }{nums: []int{2, 3, 1, 4}}, want: []int{2, 3, 4, 1}},
		{name: "test", args: struct{ nums []int }{nums: []int{2, 3, 4, 1}}, want: []int{2, 4, 1, 3}},
		{name: "test", args: struct{ nums []int }{nums: []int{2, 4, 1, 3}}, want: []int{2, 4, 3, 1}},

		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 3, 4, 5}}, want: []int{1, 2, 3, 5, 4}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 3, 5, 4}}, want: []int{1, 2, 4, 3, 5}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 4, 3, 5}}, want: []int{1, 2, 4, 5, 3}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 4, 5, 3}}, want: []int{1, 2, 5, 3, 4}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 5, 3, 4}}, want: []int{1, 2, 5, 4, 3}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 5, 4, 3}}, want: []int{1, 3, 2, 4, 5}},

		{name: "test", args: struct{ nums []int }{nums: []int{1, 1, 5}}, want: []int{1, 5, 1}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 5, 1}}, want: []int{5, 1, 1}},
		{name: "test", args: struct{ nums []int }{nums: []int{5, 1, 1}}, want: []int{1, 1, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ori := append([]int{}, tt.args.nums...)
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("Test_nextPermutation(%v) = %v, want %v", ori, tt.args.nums, tt.want)
			}
		})
	}
}
