package leetcode

// 2019/10/30 0:36 by fzls
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	n := len(nums)

	dp := make([]bool, n)
	dp[n-1] = true

	for i := n - 1; i >= 0; i-- {
		maxIdx := i + nums[i]
		if maxIdx >= n {
			maxIdx = n - 1
		}
		for j := maxIdx; j >= i+1; j-- {
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[0]
}
