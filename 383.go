package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	cntMap := make(map[int32]int)

	for _, char := range ransomNote {
		cntMap[char]++
	}

	for _, char := range magazine {
		if cntMap[char] > 0 {
			cntMap[char]--
		}
	}

	for _, cnt := range cntMap {
		if cnt != 0 {
			return false
		}
	}

	return true
}
