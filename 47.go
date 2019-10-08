package leetcode

import (
	"sort"
)

// 2019/10/09 0:36 by fzls
func permuteUnique(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)
	permuteUniqueCore(nums, 0, &res)

	return res
}

func permuteUniqueCore(nums []int, idx int, pRes *[][]int) {
	if idx == len(nums) {
		res := make([]int, len(nums))
		copy(res, nums)
		*pRes = append(*pRes, res)
		return
	}

	used := make(map[int]struct{})
	for i := idx; i < len(nums); i++ {
		if _, ok := used[nums[i]]; ok {
			continue
		}
		used[nums[i]] = struct{}{}

		nums[i], nums[idx] = nums[idx], nums[i]

		permuteUniqueCore(nums, idx+1, pRes)

		nums[i], nums[idx] = nums[idx], nums[i]
	}
}
