package leetcode

// 2019/11/17 20:55 by fzls
func subsets(nums []int) [][]int {
	sets := make([][]int, 0)
	subsetsCore(nums, 0, []int{}, &sets)
	return sets
}

func subsetsCore(nums []int, currentIdx int, set []int, pSets *[][]int) {
	if currentIdx == len(nums) {
		subSet := make([]int, len(set))
		copy(subSet, set)
		*pSets = append(*pSets, subSet)
		return
	}

	subsetsCore(nums, currentIdx+1, append(set, nums[currentIdx]), pSets)
	subsetsCore(nums, currentIdx+1, set, pSets)
}
