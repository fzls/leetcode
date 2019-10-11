package leetcode

// 2019/10/11 23:39 by fzls
func combine(n int, k int) [][]int {
	var res [][]int

	combineCore(n, k, []int{}, &res)

	return res
}

func combineCore(n int, k int, comb []int, pRes *[][]int) {
	if k == 0 {
		res := make([]int, len(comb))
		copy(res, comb)
		*pRes = append(*pRes, res)
		return
	}

	if n < k {
		return
	}

	combineCore(n-1, k, comb, pRes)
	combineCore(n-1, k-1, append(comb, n), pRes)
}
