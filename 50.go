package leetcode

// 2019/10/11 23:25 by fzls
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	powHalf := myPow(x, n/2)
	pow := powHalf * powHalf
	if n&1 != 0 {
		pow *= x
	}

	return pow
}
