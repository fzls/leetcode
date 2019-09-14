package leetcode

// 2019/09/14 17:26 by fzls
func isMatch(s string, p string) bool {
	// 如果匹配上了，则直接返回true
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	// 判断有*通配符的情况
	if len(p) >= 2 && p[1] == '*' {
		for i := 0; i <= len(s); i++ {
			// 依次匹配0至剩余最大个数个p[0]
			if i != 0 && !(p[0] == s[i-1] || p[0] == '.') {
				break
			}
			// 如果子项也匹配，则返回true
			if isMatch(s[i:], p[2:]) {
				return true
			}
		}
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
