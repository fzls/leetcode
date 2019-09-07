package leetcode

func getCntMap(s string) map[int32]int {
	cntMap := make(map[int32]int)
	for _, char := range s {
		cntMap[char]++
	}
	return cntMap
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cntMap1 := getCntMap(s)
	cntMap2 := getCntMap(t)

	if len(cntMap1) != len(cntMap2) {
		return false
	}

	for char, cnt := range cntMap1 {
		cnt2, ok := cntMap2[char]
		if !ok || cnt2 != cnt {
			return false
		}
	}

	return true
}
