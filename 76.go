package leetcode

// 2019/11/14 21:36 by fzls

func minWindow(s string, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}

	// 统计t的字符频率
	charCnts := make([]int, 256)
	for i := 0; i < len(t); i++ {
		charCnts[t[i]]++
	}

	// 滑动处理s
	l, count, minL, minR := 0, len(t), 0, len(s)
	for r := 0; r < len(s); r++ {
		// 尝试对当前字符的计数器减1
		charCnts[s[r]]--
		// 如果计数器值大于等于0，则说明这个字符是目标字符，更新count
		if charCnts[s[r]] >= 0 {
			count--
		}

		// 尝试移动左侧
		for l < r && charCnts[s[l]] < 0 {
			charCnts[s[l]]++
			l++
		}

		// 看看当前是否是一个有效地解
		if count == 0 && minR-minL > r-l {
			minL, minR = l, r
		}
	}
	if minR == len(s) {
		return ""
	}
	return s[minL : minR+1]
}
