package leetcode

type __isMatch_Key struct {
	s, p string
}

var __isMatch_cache = make(map[__isMatch_Key]bool)

// 2019/10/16 23:45 by fzls

func __isMatch(s string, p string) bool {
	return _isMatch(s, p)
}

func _isMatch(s string, p string) (matched bool) {
	if v, ok := __isMatch_cache[__isMatch_Key{s, p}]; ok {
		return v
	}
	defer func() {
		__isMatch_cache[__isMatch_Key{s, p}] = matched
	}()

	if len(p) == 0 {
		return len(s) == 0
	}

	if len(s) == 0 && p[0] != '*' {
		return false
	}

	switch p[0] {
	case '?':
		return _isMatch(s[1:], p[1:])
	case '*':
		// 0 || any
		return _isMatch(s, p[1:]) || len(s) > 0 && _isMatch(s[1:], p)
	default:
		return s[0] == p[0] && _isMatch(s[1:], p[1:])
	}
}
