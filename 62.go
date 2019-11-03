package leetcode

// 2019/11/03 20:23 by fzls
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	// init dp
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	// at row i: dp[j] = dp[j](last row) + dp[j+1](this row)
	for row := m - 2; row >= 0; row-- {
		for col := n - 2; col >= 0; col-- {
			dp[col] += dp[col+1]
		}
	}

	return dp[0]
}
