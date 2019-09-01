package leetcode

func isCharOrDigit(c byte) bool {
	// 是否是数字或字母
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func isSame(c1, c2 byte) bool {
	// 一样
	if c1 == c2 {
		return true
	}

	// 大小写不同也视为一样
	if c1 >= 'a' && c1 <= 'z' && c1-c2 == 'a'-'A' ||
		c2 >= 'a' && c2 <= 'z' && c2-c1 == 'a'-'A' {
		return true
	}

	return false
}

func isPalindrome_(s string) bool {
	chars := []byte(s)

	i := 0
	j := len(chars) - 1
	for i < j {
		// 跳过所有非字母和数字的字符
		for i <= j && !isCharOrDigit(chars[i]) {
			i++
		}
		for i <= j && !isCharOrDigit(chars[j]) {
			j--
		}

		// 没有其他字符需要处理了
		if i > j {
			return true
		}

		// 当前字符不同则直接返回否
		if !isSame(chars[i], chars[j]) {
			return false
		}

		i++
		j--
	}

	return true
}
