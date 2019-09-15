package leetcode

type __key struct {
	s, p string
}

var cache = make(map[__key]bool)

// 2019/09/14 17:26 by fzls
func isMatch(s string, p string) (res bool) {
	// 对重复计算进行缓存
	if matched, ok := cache[__key{s, p}]; ok {
		return matched
	}

	defer func() {
		cache[__key{s, p}] = res
	}()

	// 如果pattern已全部消耗，则只需判断字符串是否也已消耗完
	if len(p) == 0 {
		return len(s) == 0
	}

	// 判断第一个字符是否匹配
	firstMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		// 判断有*通配符的情况
		return isMatch(s, p[2:]) || // 不匹配*
			(firstMatch && isMatch(s[1:], p)) // 消耗掉一个源字符串的字符
	} else {
		// 无通配符
		return firstMatch && isMatch(s[1:], p[1:])
	}
}
