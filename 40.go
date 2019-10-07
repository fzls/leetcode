package leetcode

import (
	"sort"
)

// 2019/10/07 23:07 by fzls
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var combs [][]int
	combinationSum2Core(candidates, target, 0, false, []int{}, &combs)
	return combs
}

func combinationSum2Core(candidates []int, target int, idx int, useLast bool, comb []int, pCombs *[][]int) {
	if target < 0 {
		return
	}

	if idx == len(candidates) {
		if target == 0 {
			*pCombs = append(*pCombs, append([]int{}, comb...))
		}
		return
	}

	// use it
	if !(!useLast && idx != 0 && candidates[idx-1] == candidates[idx]) {
		combinationSum2Core(candidates, target-candidates[idx], idx+1, true, append(comb, candidates[idx]), pCombs)
	}

	// not use it
	combinationSum2Core(candidates, target, idx+1, false, comb, pCombs)
}
