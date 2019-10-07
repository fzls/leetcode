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
			res := make([]int, len(comb))
			copy(res, comb)
			*pCombs = append(*pCombs, res)
		}
		return
	}

	// use it
	// 当之前的一个相同的数字没有使用的时候，不遍历本数字采用的情况，因为会造成重复
	// 比如假如是 1,1,1,4；现在处理到第三个1的时候，假如第二个没有采用，在其之前的组合为comb，若遍历本次的1，则由comb,1
	// 但是，考虑到第二个1采用，本次的1没有采用的情况，这两种会造成重复，所以我们在这种情况只保留前面采用了，而后面一个不采用这种组合
	if !(!useLast && idx != 0 && candidates[idx-1] == candidates[idx]) {
		combinationSum2Core(candidates, target-candidates[idx], idx+1, true, append(comb, candidates[idx]), pCombs)
	}

	// not use it
	combinationSum2Core(candidates, target, idx+1, false, comb, pCombs)
}
