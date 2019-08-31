package leetcode

func mySqrt(x int) int {
	low := 1
	high := x

	for low <= high {
		i := low + (high-low)/2

		res := i * i
		if res == x {
			return i
		} else if res < x {
			low = i + 1
		} else {
			high = i - 1
		}
	}

	return high
}
