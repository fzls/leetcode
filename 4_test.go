package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		numsA []int
		numsB []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		struct {
			name string
			args args
			want float64
		}{name: string("A"), args: struct {
			numsA []int
			numsB []int
		}{numsA: []int{1, 3}, numsB: []int{2}}, want: 2.0},
		struct {
			name string
			args args
			want float64
		}{name: string("B"), args: struct {
			numsA []int
			numsB []int
		}{numsA: []int{1, 2}, numsB: []int{3, 4}}, want: 2.5},
		struct {
			name string
			args args
			want float64
		}{name: string("C"), args: struct {
			numsA []int
			numsB []int
		}{numsA: []int{1, 2, 3}, numsB: []int{3, 4}}, want: 3.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _findMedianSortedArrays(tt.args.numsA, tt.args.numsB); got != tt.want {
				t.Errorf("_findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
