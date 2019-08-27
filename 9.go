package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	n := 0
	for _x := x; _x > 0; _x /= 10 {
		n++
	}

	// 10^(n-1)
	leftRatio := 1
	for i := 0; i < n-1; i++ {
		leftRatio *= 10
	}
	// 1
	rightRatio := 1

	for i := 1; 2*i <= n; i++ {
		left := x / leftRatio % 10
		right := x / rightRatio % 10
		if left != right {
			return false
		}

		leftRatio /= 10
		rightRatio *= 10
	}

	return true
}
