package leetcode

// 2019/09/14 17:26 by fzls
func isMatch(s string, p string) bool {
	// 如果匹配上了，则直接返回true
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	// 判断有*通配符的情况
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) ||
			(p[0] == s[0] || p[0] == '.') && isMatch(s[1:], p)
	}

	// 没有通配符的情况下，如果此时两者有一个为空，则说明不匹配
	if len(s) == 0 || len(p) == 0 {
		return false
	}

	// 否则看首字符是否匹配
	if p[0] == s[0] || p[0] == '.' {
		return isMatch(s[1:], p[1:])
	} else {
		return false
	}
}
