package leetcode

import (
	"fmt"
	"sort"
)

// 2019/10/07 23:07 by fzls
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var combs [][]int
	combinationSum2Core(candidates, target, 0, []int{}, &combs, make(map[string]struct{}))
	return combs
}

func sameList(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func combinationSum2Core(candidates []int, target int, idx int, comb []int, pCombs *[][]int, set map[string]struct{}) {
	if target < 0 {
		return
	}

	if idx == len(candidates) {
		if target == 0 {
			key := fmt.Sprintf("%v", comb)
			if _, ok := set[key]; ok {
				return
			}
			set[key] = struct{}{}
			*pCombs = append(*pCombs, append([]int{}, comb...))
		}
		return
	}

	// use it
	combinationSum2Core(candidates, target-candidates[idx], idx+1, append(comb, candidates[idx]), pCombs, set)

	// not use it
	combinationSum2Core(candidates, target, idx+1, comb, pCombs, set)
}
