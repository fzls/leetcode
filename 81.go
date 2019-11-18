package leetcode

// 2019/11/18 22:20 by fzls
func __search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) >> 1

		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			if nums[mid] < nums[high] && nums[high] < target {
				high = mid - 1
			}
			if nums[mid] > nums[low] {
				low = mid + 1
			}
		} else { // nums[mid] > target
			if nums[low] < nums[mid] && nums[low] > target {
				low = mid + 1
			}
			if nums[mid] < nums[high] {
				high = mid - 1
			}
		}
		if high >= 0 && nums[high] == target {
			return true
		} else {
			high--
		}
		if low < len(nums) && nums[low] == target {
			return true
		} else {
			low++
		}
	}

	return false
}
