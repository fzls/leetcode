package leetcode

func firstUniqChar(s string) int {
	charCnt := make([]int, 26)
	for _, char := range s {
		charCnt[char-'a']++
	}

	for idx, char := range s {
		if charCnt[char-'a'] == 1 {
			return idx
		}
	}

	return -1
}
