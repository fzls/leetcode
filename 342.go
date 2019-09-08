package leetcode

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}

	for num&3 == 0 {
		num >>= 2
	}

	return num == 1
}
