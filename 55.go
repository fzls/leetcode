package leetcode

// 2019/10/30 0:36 by fzls
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	n := len(nums)

	dp := make([]bool, n)
	dp[n-1] = true
	lastCanJump := n - 1

	for i := n - 1; i >= 0; i-- {
		if i+nums[i] < lastCanJump {
			continue
		}

		dp[i] = true
		lastCanJump = i
	}

	return dp[0]
}
