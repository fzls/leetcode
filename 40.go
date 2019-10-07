package leetcode

import (
	"reflect"
	"sort"
)

// 2019/10/07 23:07 by fzls
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var combs [][]int
	combinationSum2Core(candidates, target, 0, []int{}, &combs)
	return combs
}

func combinationSum2Core(candidates []int, target int, idx int, comb []int, pCombs *[][]int) {
	if target < 0 {
		return
	}

	if idx == len(candidates) {
		if target == 0 {
			for _, cb := range *pCombs {
				if reflect.DeepEqual(comb, cb) {
					return
				}
			}
			*pCombs = append(*pCombs, append([]int{}, comb...))
		}
		return
	}

	// use it
	combinationSum2Core(candidates, target-candidates[idx], idx+1, append(comb, candidates[idx]), pCombs)

	// not use it
	combinationSum2Core(candidates, target, idx+1, comb, pCombs)
}
