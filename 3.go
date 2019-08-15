package leetcode

func lengthOfLongestSubstring(s string) int {
	var maxLen int
	for i := 0; i < len(s) && i+maxLen < len(s); i++ {
		set := make(map[byte]struct{})
		for j := i; j < len(s); j++ {
			if _, exists := set[s[j]]; !exists {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
				}
				set[s[j]] = struct{}{}
			} else {
				break
			}
		}
	}

	return maxLen
}
