package leetcode

func lengthOfLongestSubstring(s string) int {
	var maxLen, i, j int
	// 滑动窗口
	mapChar2Idx := make(map[byte]int)
	for i < len(s) && j < len(s) && i+maxLen < len(s) {
		if idx, exists := mapChar2Idx[s[j]]; !exists {
			// 若此字符在之前未出现，则尝试更新最大值，并记录位置
			if j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		} else {
			for _i := i; _i <= idx; _i++ {
				delete(mapChar2Idx, s[_i])
			}
			// 否则直接将左侧下标推进至与之前出现的本字符的后面一位
			i = idx + 1
		}
		mapChar2Idx[s[j]] = j
		j++
	}

	return maxLen
}
