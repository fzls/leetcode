package leetcode

// 2019/09/28 23:03 by fzls

func search(nums []int, target int) int {
	return _search2(nums, target)
}

func _search(nums []int, target int) int {
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
	// 对多种情况进行分类讨论
	for low <= high {
		i := int(uint(low+high) >> 1)
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			if nums[low] <= nums[i] {
				// nums[low] <= nums[i] < target
				low = i + 1
			} else if nums[i] <= nums[high] {
				if nums[high] < target {
					// nums[i] <= nums[high] < target
					high = i - 1
				} else {
					// nums[i] < target <= nums[high]
					return __bs(nums, i+1, high, target)
				}
			}
		} else {
			if nums[i] <= nums[high] {
				// target < nums[i] <= nums[high]
				high = i - 1
			} else if nums[low] <= nums[i] {
				if target < nums[low] {
					// target < nums[low] <= nums[i]
					low = i + 1
				} else {
					// nums[low] <= target <= nums[i]
					return __bs(nums, low, i-1, target)
				}
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

func _search2(nums []int, target int) int {
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

	changePoint := getChangePoint(nums)

	if changePoint == -1 {
		return __bs(nums, 0, len(nums)-1, target)
	} else {
		if target >= nums[0] {
			return __bs(nums, 0, changePoint, target)
		} else {
			return __bs(nums, changePoint+1, len(nums)-1, target)
		}
	}
}

func getChangePoint(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return -1
	}

	low := 0
	high := len(nums) - 1
	for low <= high {
		i := int(uint(low+high) >> 1)
		if nums[i] > nums[i+1] {
			return i
		} else {
			if nums[i] >= nums[0] {
				low = i + 1
			} else {
				high = i - 1
			}
		}
	}

	return 0
}
