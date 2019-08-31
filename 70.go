package leetcode

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	before2 := 1
	before1 := 2
	current := 0

	for i := 2; i < n; i++ {
		current = before1 + before2
		before2 = before1
		before1 = current
	}

	return current
}
