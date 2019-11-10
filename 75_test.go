package leetcode

import (
	"reflect"
	"testing"
)

// 2019/11/11 0:01 by fzls

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{2, 0, 2, 1, 1, 0}}, want: []int{0, 0, 1, 1, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.want, tt.args.nums) {
				t.Errorf("sortColors() \nres= %v, \nwant %v", tt.args, tt.want)
			}
		})
	}
}
