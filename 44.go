package leetcode

// 2019/10/16 23:45 by fzls
func __isMatch(s string, p string) bool {
	return _matchDp(s, p)
	return newIsMatchSolution(s, p).solve()
}

func _matchDp(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}

type __isMatch_Key struct {
	sIdx, pIdx int
}

type IsMatchSolution struct {
	s, p []byte

	cache map[__isMatch_Key]bool
}

func newIsMatchSolution(s, p string) *IsMatchSolution {
	return &IsMatchSolution{
		s:     []byte(s),
		p:     []byte(p),
		cache: make(map[__isMatch_Key]bool),
	}
}
func (so *IsMatchSolution) solve() bool {
	return so.solveCore(0, 0)
}

func (so *IsMatchSolution) solveCore(sIdx, pIdx int) (matched bool) {
	if v, ok := so.cache[__isMatch_Key{sIdx, pIdx}]; ok {
		return v
	}
	defer func() {
		so.cache[__isMatch_Key{sIdx, pIdx}] = matched
	}()

	if pIdx == len(so.p) {
		return sIdx == len(so.s)
	}

	if sIdx == len(so.s) && so.p[pIdx] != '*' {
		return false
	}

	switch so.p[pIdx] {
	case '?':
		return so.solveCore(sIdx+1, pIdx+1)
	case '*':
		// 0 || any
		return so.solveCore(sIdx, pIdx+1) || sIdx+1 <= len(so.s) && so.solveCore(sIdx+1, pIdx)
	default:
		return so.s[sIdx] == so.p[pIdx] && so.solveCore(sIdx+1, pIdx+1)
	}
}
