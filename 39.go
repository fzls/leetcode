package leetcode

// 2019/10/04 20:30 by fzls
func combinationSum(candidates []int, target int) [][]int {
	var combs [][]int
	combinationSumCore(candidates, len(candidates)-1, target, []int{}, &combs)
	return combs
}

func combinationSumCore(candidates []int, idx, target int, comb []int, combs *[][]int) {
	if target < 0 {
		return
	}

	if idx == -1 {
		if target == 0 {
			*combs = append(*combs, append([]int{}, comb...))
		}

		return
	}

	for _target := target; _target >= 0; _target -= candidates[idx] {
		combinationSumCore(candidates, idx-1, _target, comb, combs)
		comb = append(comb, candidates[idx])
	}
}
