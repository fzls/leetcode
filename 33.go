package leetcode

// 2019/09/28 23:03 by fzls
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	if nums[0] < nums[len(nums)-1] {
		// 完全排序的情况，直接使用二分查找
		return __bs(nums, 0, len(nums)-1, target)
	}

	// 其他情况
	low := 0
	high := len(nums) - 1
	for low <= high {
		i := int(uint(low+high) >> 1)
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			if nums[i] >= nums[low] {
				low = i + 1
			} else if nums[i] <= nums[high] {
				if nums[high] < target {
					high = i - 1
				} else {
					return __bs(nums, i+1, high, target)
				}
			}
		} else {
			// nums[i]>target
			if nums[i] >= nums[low] {
				if nums[low] > target {
					low = i + 1
				} else {
					return __bs(nums, low, i-1, target)
				}
			} else if nums[i] <= nums[high] {
				high = i - 1
			}
		}
	}

	return -1
}

func __bs(nums []int, low, high int, target int) int {
	for low <= high {
		i := int(uint(low+high) >> 1)
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			low = i + 1
		} else {
			high = i - 1
		}
	}
	return -1
}
