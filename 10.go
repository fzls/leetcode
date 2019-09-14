package leetcode

// 2019/09/14 17:26 by fzls
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	if len(p) >= 2 && p[1] == '*' {
		for i := 0; i <= len(s); i++ {
			if i != 0 && !(p[0] == s[i-1] || p[0] == '.') {
				break
			}
			if isMatch(s[i:], p[2:]) {
				return true
			}
		}
	}

	if len(s) == 0 || len(p) == 0 {
		return false
	}

	if p[0] == s[0] || p[0] == '.' {
		return isMatch(s[1:], p[1:])
	} else {
		return false
	}
}
