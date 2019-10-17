package leetcode

import (
	"math"
)

// 2019/10/17 23:56 by fzls
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = 0

	for i := n - 2; i >= 0; i-- {
		if i+nums[i] >= n-1 {
			dp[i] = 1
			continue
		}

		ms := math.MaxInt32
		for j := i + 1; j <= i+nums[i] && j < n; j++ {
			if dp[j] < ms {
				ms = dp[j]
			}
		}
		dp[i] = 1 + ms
	}

	return dp[0]
}
