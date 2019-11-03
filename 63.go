package leetcode

// 2019/11/03 20:50 by fzls
const HAS_STONE = 1

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	// special case
	if obstacleGrid[0][0] == HAS_STONE || obstacleGrid[m-1][n-1] == HAS_STONE {
		return 0
	}

	// dp(reuse memory)
	dp := obstacleGrid
	// init
	dp[m-1][n-1] = 1
	for row := m - 2; row >= 0; row-- {
		if obstacleGrid[row][n-1] == HAS_STONE {
			dp[row][n-1] = 0
			continue
		}
		dp[row][n-1] = dp[row+1][n-1]
	}
	for col := n - 2; col >= 0; col-- {
		if obstacleGrid[m-1][col] == HAS_STONE {
			dp[m-1][col] = 0
			continue
		}
		dp[m-1][col] = dp[m-1][col+1]
	}

	// iter
	for row := m - 2; row >= 0; row-- {
		for col := n - 2; col >= 0; col-- {
			if obstacleGrid[row][col] == HAS_STONE {
				dp[row][col] = 0
				continue
			}
			dp[row][col] = dp[row+1][col] + dp[row][col+1]
		}
	}

	return dp[0][0]
}
