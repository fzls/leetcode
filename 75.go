package leetcode

// 2019/11/10 23:57 by fzls
func sortColors(nums []int) {
	cur := 0
	redHigh := 0
	blueLow := len(nums) - 1

	for cur <= blueLow {
		if nums[cur] == 0 {
			nums[cur], nums[redHigh] = nums[redHigh], nums[cur]
			redHigh++
			cur++
		} else if nums[cur] == 2 {
			nums[cur], nums[blueLow] = nums[blueLow], nums[cur]
			blueLow--
		} else {
			cur++
		}
	}

	return
}
