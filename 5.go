package leetcode

func longestPalindrome(s string) string {
	return ManachersAlgorithm(s)
}

// 阅读博文后实现的O(n)算法
// https://medium.com/hackernoon/manachers-algorithm-explained-longest-palindromic-substring-22cb27a5e96f
func ManachersAlgorithm(s string) string {
	if len(s) <= 1{
		return s
	}

	// 统一插上特殊字符#，使得字符串总为奇数长
	s = getModifiedStr(s)

	// 以下标i为中心的最大回文子串的半径
	p := make([]int, len(s))
	// 到当前为止的回文子串的中心及右边缘
	c := 0
	r := 0
	// 记录到目前遇到的最大的回文子串的中心和半径
	maxC := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if i < r {
			// 如果当前仍在上个回文子串的半径内，则初始扩展值为min(p[mirror], r-i)
			p[i] = p[2*c-i]
			if p[i] > r-i {
				p[i] = r - i
			}
		}
		// 以i为中心，继续扩展半径
		a := i + p[i] + 1
		b := i - p[i] - 1
		for a < len(s) && b >= 0 && s[a] == s[b] {
			p[i]++
			a++
			b--
		}

		// 如果i的回文子串越出了之前的回文子串的范围，则更新中心和右边缘，并判断是否比目前记录的最大值要大
		if i+p[i] > r {
			c = i
			r = i + p[i]
			if p[i] > maxLen {
				maxC = i
				maxLen = p[i]
			}
		}
	}

	// 将结果范围去除特殊字符后返回
	return getUnmodifiedStr(s[maxC-maxLen:maxC+maxLen+1])
}

const SpecialChar = '#'

func getModifiedStr(s string) string {
	res := make([]byte, 2*len(s)+1)
	res[0] = SpecialChar
	for idx, char := range []byte(s) {
		res[2*idx+1] = char
		res[2*idx+2] = SpecialChar
	}
	return string(res)
}

func getUnmodifiedStr(s string) string {
	var res []byte
	for _, char := range []byte(s){
		if char != SpecialChar {
			res = append(res, char)
		}
	}
	return string(res)
}
