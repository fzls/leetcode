package leetcode

// 2019/11/03 21:10 by fzls
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	// init
	dp := grid
	for row := m - 2; row >= 0; row-- {
		dp[row][n-1] += dp[row+1][n-1]
	}
	for col := n - 2; col >= 0; col-- {
		dp[m-1][col] += dp[m-1][col+1]
	}

	// iter: dp[r][c] = grid[r][c] + min(dp[r+1][c], dp[r][c+1])
	for row := m - 2; row >= 0; row-- {
		for col := n - 2; col >= 0; col-- {
			next := dp[row+1][col]
			if dp[row][col+1] < next {
				next = dp[row][col+1]
			}

			dp[row][col] += next
		}
	}

	return dp[0][0]
}
