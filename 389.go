package leetcode

func findTheDifference(s string, t string) byte {
	cntMap := make(map[byte]int)
	for _, c := range []byte(t) {
		cntMap[c]++
	}

	for _, c := range []byte(s) {
		cntMap[c]--
	}

	for c, cnt := range cntMap {
		if cnt != 0 {
			return c
		}
	}

	return '0'
}
