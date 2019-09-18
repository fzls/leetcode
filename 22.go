package leetcode

import "fmt"

// 2019/09/18 22:58 by fzls
func generateParenthesis(n int) []string {
	return _dp_generateParenthesis(n)

	var res []string
	_generateParenthesis(0, 2*n, []byte{}, &res)
	return res
}

func _generateParenthesis(nUnpaired int, nRemaining int, temp []byte, res *[]string) {
	// 如果剩余括号数小于未匹配数，则不合法，无需继续查找
	if nRemaining < nUnpaired {
		return
	}

	// 找到一个组合
	if nRemaining == 0 {
		*res = append(*res, string(temp))
		return
	}

	// 搜索
	if nRemaining > nUnpaired {
		// 剩余括号数大于未匹配数，则说明本次可以尝试放左括号
		_generateParenthesis(nUnpaired+1, nRemaining-1, append(temp, '('), res)
	}
	if nUnpaired > 0 {
		// 未匹配数不为0， 则说明本次可以尝试放右括号
		_generateParenthesis(nUnpaired-1, nRemaining-1, append(temp, ')'), res)
	}
}

// 参考官方题解的动态规划方法
// 因为第一个位置必定为左括号，我们可以通过枚举其匹配的右括号的所有可能位置，将问题转化为子问题
// 假设dp(n)是n个左右括号的所有组合，右括号可能在1,3,5,...2*i+1,...2*(n-1)+1 的位置，假设为位置为2*i+1，
// 则这种情况下所有可能可以表示为 ( + left + ) + right，其中left=dp(i), right = dp(n-i-1)
func _dp_generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	var res []string
	for i := 0; i < n; i++ {
		for _, left := range _dp_generateParenthesis(i) {
			for _, right := range _dp_generateParenthesis(n - i - 1) {
				res = append(res, fmt.Sprintf("(%v)%v", left, right))
			}
		}
	}
	return res
}
