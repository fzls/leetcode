package leetcode

import "math"

// 2019/10/17 23:56 by fzls
func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	dp := make([]int, n)
	dp[n-1] = 0

	for i := n - 2; i >= 0; i-- {
		if i+nums[i] >= n-1 {
			dp[i] = 1
			continue
		}

		maxIdx := i + nums[i]
		if maxIdx >= n {
			maxIdx = n - 1
		}
		ms := math.MaxInt32
		for j := maxIdx; j > i; j-- {
			if dp[j] < ms {
				ms = dp[j]
				if ms == 1 {
					break
				}
			}
		}
		dp[i] = 1 + ms
	}

	return dp[0]
}
