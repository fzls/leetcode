package leetcode

// 2019/11/07 21:23 by fzls
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// alloc
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// init
	for i := 0; i <= m; i++ {
		dp[i][n] = m - i
	}
	for j := 0; j <= n; j++ {
		dp[m][j] = n - j
	}

	// word1[i] == word2[j] => dp[i][j] = d[i+1][j+1]
	// word1[i] != word2[j] => dp[i][j] = 1 + min( dp[i][j+1], dp[i+1][j], dp[i+1][j+1] )
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				// insert
				min := dp[i][j+1]
				// delete
				if dp[i+1][j] < min {
					min = dp[i+1][j]
				}
				// replace
				if dp[i+1][j+1] < min {
					min = dp[i+1][j+1]
				}

				// op+1
				dp[i][j] = 1 + min
			}
		}
	}

	return dp[0][0]
}
