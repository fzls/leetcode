package leetcode

// 2019/11/07 21:23 by fzls
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// alloc
	dp := make([][]int, 2)
	for i := 0; i <= 1; i++ {
		dp[i] = make([]int, n+1)
	}

	// init
	currentIdx := 0
	nextIdx := 1
	current := dp[currentIdx]
	next := dp[nextIdx]
	for j := 0; j <= n; j++ {
		current[j] = n - j
	}

	for i := m - 1; i >= 0; i-- {
		currentIdx, nextIdx = 1-currentIdx, 1-nextIdx
		current, next = dp[currentIdx], dp[nextIdx]

		current[n] = 1 + next[n]

		for j := n - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				current[j] = next[j+1]
			} else {
				// insert
				min := current[j+1]
				// delete
				if next[j] < min {
					min = next[j]
				}
				// replace
				if next[j+1] < min {
					min = next[j+1]
				}

				// op+1
				current[j] = 1 + min
			}
		}
	}

	return current[0]
}
