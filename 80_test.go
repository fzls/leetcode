package leetcode

import (
	"reflect"
	"testing"
)

// 2019/11/18 22:11 by fzls

func Test__removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{1, 1, 1, 2, 2, 3}}, want: []int{1, 1, 2, 2, 3}},
		{name: "test", args: struct{ nums []int }{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, want: []int{0, 0, 1, 1, 2, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := _removeDuplicates(tt.args.nums)
			if !(got == len(tt.want) && reflect.DeepEqual(tt.args.nums[:got], tt.want)) {
				t.Errorf("_removeDuplicates(%v) = %v, want %v", tt.args.nums, tt.args.nums[:got], tt.want)
			}
		})
	}
}
