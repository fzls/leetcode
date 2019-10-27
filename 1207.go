package leetcode

// 2019/10/27 20:36 by fzls
func uniqueOccurrences(arr []int) bool {
	numberCounts := make(map[int]int) // number -> counts
	for _, number := range arr {
		numberCounts[number]++
	}

	countSet := make(map[int]struct{})
	for _, cnt := range numberCounts {
		if _, exists := countSet[cnt]; exists {
			return false
		}
		countSet[cnt] = struct{}{}
	}

	return true
}
