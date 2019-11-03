package leetcode

// 2019/11/03 20:23 by fzls
func uniquePaths(m int, n int) int {
	// todo: check if can use less memory
	// init dp
	dp := make([][]int, m)
	buf := make([]int, m*n)
	for i := 0; i < m; i++ {
		dp[i], buf = buf[:n], buf[n:]
	}

	// dp[i][j] = dp[i+1][j] + dp[i][j+1]
	for row := 0; row < m; row++ {
		dp[row][n-1] = 1
	}
	for col := 0; col < n; col++ {
		dp[m-1][col] = 1
	}
	for row := m - 2; row >= 0; row-- {
		for col := n - 2; col >= 0; col-- {
			dp[row][col] = dp[row+1][col] + dp[row][col+1]
		}
	}

	return dp[0][0]
}
