package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var res []byte

	minLen := 0x7FFFFFFF
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

OUTER:
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[j-1][i] {
				break OUTER
			}
		}
		res = append(res, strs[0][i])
	}

	return string(res)
}
