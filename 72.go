package leetcode

// 2019/11/07 21:23 by fzls
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// alloc
	dp := make([]int, n+1)

	// init
	for j := 0; j <= n; j++ {
		dp[j] = n - j
	}

	// to save space, wo use only one row, but we need to save dp[i+1][j+1] into nextLineNextRowValue before we try to change dp[i+1][j+1]
	// word1[i] == word2[j] => dp[i][j] = d[i+1][j+1]
	// word1[i] != word2[j] => dp[i][j] = 1 + min( dp[i][j+1], dp[i+1][j], dp[i+1][j+1] )
	for i := m - 1; i >= 0; i-- {
		nextLineNextRowValue := dp[n] // dp[i+1][j+1]
		dp[n]++

		for j := n - 1; j >= 0; j-- {
			temp := dp[j]
			if word1[i] == word2[j] {
				dp[j] = nextLineNextRowValue
			} else {
				// insert
				min := dp[j+1]
				// delete
				if dp[j] < min {
					min = dp[j]
				}
				// replace
				if nextLineNextRowValue < min {
					min = nextLineNextRowValue
				}

				// op+1
				dp[j] = 1 + min
			}
			nextLineNextRowValue = temp
		}
	}

	return dp[0]
}
