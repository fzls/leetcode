package leetcode

// 2019/10/12 22:59 by fzls
func distributeCandies(candies []int) int {
	n := len(candies)

	set := make(map[int]struct{})
	for _, c := range candies {
		set[c] = struct{}{}
	}

	maxCount := len(set)
	if maxCount >= n/2 {
		maxCount = n / 2
	}

	return maxCount
}
