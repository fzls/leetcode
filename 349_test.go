package leetcode

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", args: struct {
			nums1 []int
			nums2 []int
		}{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}, want: []int{2}},
		{name: "test", args: struct {
			nums1 []int
			nums2 []int
		}{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}, want: []int{9, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(listIntToSet(got), listIntToSet(tt.want)) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
