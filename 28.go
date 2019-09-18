package leetcode

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i+len(needle)-1 < len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
