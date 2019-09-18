package leetcode

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	low := 2
	high := num / 2
	for low <= high {
		i := low + (high-low)/2
		pow2 := i * i
		if pow2 == num {
			return true
		} else if pow2 < num {
			low = i + 1
		} else {
			high = i - 1
		}
	}

	return false
}
