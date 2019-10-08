package leetcode

// 2019/10/09 0:26 by fzls
func permute(nums []int) [][]int {
	var res [][]int

	permuteCore(nums, 0, &res)

	return res
}

func permuteCore(nums []int, idx int, pRes *[][]int) {
	if idx == len(nums) {
		res := make([]int, len(nums))
		copy(res, nums)
		*pRes = append(*pRes, res)
		return
	}

	for i := idx; i < len(nums); i++ {
		// swap
		nums[idx], nums[i] = nums[i], nums[idx]

		// per
		permuteCore(nums, idx+1, pRes)

		// swap back
		nums[idx], nums[i] = nums[i], nums[idx]
	}
}
