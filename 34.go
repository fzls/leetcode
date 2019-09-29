package leetcode

// 2019/09/29 21:51 by fzls
func searchRange(nums []int, target int) []int {
	return []int{
		searchFirst(nums, target),
		searchLast(nums, target),
	}
}

func searchFirst(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := int(uint(low+high) / 2)

		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == 0 ||
				mid > 0 && nums[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func searchLast(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := int(uint(low+high) / 2)

		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == len(nums)-1 ||
				mid < len(nums)-1 && nums[mid+1] > target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}
