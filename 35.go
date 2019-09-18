package leetcode

// 二分查找
func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		i := low + (high-low)/2
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			low = i + 1
		} else {
			high = i - 1
		}
	}

	return low
}
