package leetcode

func longestPalindrome_(s string) int {
	cntMap := make(map[int32]int)

	for _, char := range s {
		cntMap[char]++
	}

	maxLen := 0
	hasOdd := false
	for _, cnt := range cntMap {
		if cnt%2 == 0 {
			maxLen += cnt
		} else {
			maxLen += cnt - 1
			hasOdd = true
		}
	}

	if hasOdd {
		maxLen++
	}
	return maxLen
}
