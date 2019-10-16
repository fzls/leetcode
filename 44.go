package leetcode

// 2019/10/16 23:45 by fzls

func __isMatch(s string, p string) bool {
	return newIsMatchSolution(s, p).solve()
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
