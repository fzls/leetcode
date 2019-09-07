package leetcode

func _swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func moveZeroes(nums []int) {
	zeroIdx := 0
	nonZeroIdx := 0
	for nonZeroIdx < len(nums) {
		for zeroIdx < len(nums) && nums[zeroIdx] != 0 {
			zeroIdx++
		}
		if nonZeroIdx < zeroIdx {
			nonZeroIdx = zeroIdx + 1
		} else {
			nonZeroIdx++
		}
		for nonZeroIdx < len(nums) && nums[nonZeroIdx] == 0 {
			nonZeroIdx++
		}
		if nonZeroIdx < len(nums) {
			_swap(nums, zeroIdx, nonZeroIdx)
		}
	}
}
