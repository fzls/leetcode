package leetcode

func lengthOfLongestSubstring(s string) int {
	var maxLen int
	// 依次遍历子字符串左侧下标
	for i := 0; i < len(s) && i+maxLen < len(s); {
		// 记录从这里开始的已出现的字符，及其位置
		mapChar2Idx := make(map[byte]int)
		for j := i; j < len(s); j++ {
			if idx, exists := mapChar2Idx[s[j]]; !exists {
				// 若此字符在之前未出现，则尝试更新最大值，并记录位置
				if j-i+1 > maxLen {
					maxLen = j - i + 1
				}
				mapChar2Idx[s[j]] = j
			} else {
				// 否则直接将左侧下标推进至与之前出现的本字符的后面一位
				i = idx + 1
				break
			}
		}
	}

	return maxLen
}
