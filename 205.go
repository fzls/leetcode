package leetcode

func getIndexes(s string) []int {
	firstAppearIndex := make([]int, 256)
	idx := 1
	var indexes []int
	for _, char := range s {
		if firstAppearIndex[char] == 0 {
			firstAppearIndex[char] = idx // char第一次出现为第idx个（1-n)
			idx++
		}
		indexes = append(indexes, firstAppearIndex[char])
	}
	return indexes
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) <= 1 {
		return true
	}

	// 假设某个字符char在字符串中首次出现的下标为idx[char]
	// 将字符串每个字母替换为其首次出现的下标，则构成一个下标数组
	// 两个字符串同构等价于下标数组一致
	indexes1 := getIndexes(s)
	indexes2 := getIndexes(t)
	for i := 0; i < len(s); i++ {
		if indexes1[i] != indexes2[i] {
			return false
		}
	}

	return true
}
