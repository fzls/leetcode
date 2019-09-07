package leetcode

import "strings"

func getPatternByte(data []byte) []int {
	var pattern []int

	cnt := 0
	cntMap := make(map[byte]int)
	for _, v := range data {
		if _cnt, ok := cntMap[v]; ok {
			pattern = append(pattern, _cnt)
		} else {
			cntMap[v] = cnt
			pattern = append(pattern, cnt)
			cnt++
		}
	}

	return pattern
}

func getPatternString(data []string) []int {
	var pattern []int

	cnt := 0
	cntMap := make(map[string]int)
	for _, v := range data {
		if _cnt, ok := cntMap[v]; ok {
			pattern = append(pattern, _cnt)
		} else {
			cntMap[v] = cnt
			pattern = append(pattern, cnt)
			cnt++
		}
	}

	return pattern
}

func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}

	pattern1 := getPatternByte([]byte(pattern))
	pattern2 := getPatternString(words)

	for i := 0; i < len(pattern); i++ {
		if pattern1[i] != pattern2[i] {
			return false
		}
	}

	return true
}
