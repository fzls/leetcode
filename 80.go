package leetcode

// 2019/11/18 21:57 by fzls
func _removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	// place to put next item(one past to the valid items)
	left := 0
	// counts for the current number
	currentCounts := 0
	lastNum := nums[0]
	const maxDup = 2
	// place search so far
	for right := 0; right < len(nums); right++ {
		if nums[right] != lastNum {
			currentCounts = 0
			lastNum = nums[right]
		}
		currentCounts++
		if currentCounts <= maxDup {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	return left
}
