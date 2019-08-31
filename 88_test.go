package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	{
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		t.Run("test", func(t *testing.T) {
			merge(nums1, 3, nums2, 3)
			if !reflect.DeepEqual(nums1, []int{1, 2, 2, 3, 5, 6}) {
				t.Errorf("merge() = %v, want %v", nums1, []int{1, 2, 2, 3, 5, 6})
			}
		})
	}
	{
		nums1 := []int{4, 5, 6, 0, 0, 0}
		nums2 := []int{1, 2, 3}
		t.Run("test", func(t *testing.T) {
			merge(nums1, 3, nums2, 3)
			if !reflect.DeepEqual(nums1, []int{1, 2, 3, 4, 5, 6}) {
				t.Errorf("merge() = %v, want %v", nums1, []int{1, 2, 3, 4, 5, 6})
			}
		})
	}
}
