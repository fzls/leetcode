package leetcode

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	stairs := make([]int, n)
	stairs[0] = 1
	stairs[1] = 2

	for i := 2; i < n; i++ {
		stairs[i] = stairs[i-1] + stairs[i-2]
	}

	return stairs[n-1]
}
