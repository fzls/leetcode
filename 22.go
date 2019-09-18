package leetcode

// 2019/09/18 22:58 by fzls
func generateParenthesis(n int) []string {
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
