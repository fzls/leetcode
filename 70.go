package leetcode

var cache = make(map[int]int)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	if c, ok := cache[n]; ok {
		return c
	}

	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]
}
