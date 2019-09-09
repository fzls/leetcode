package leetcode

func findNthDigit(n int) int {
	l := 1
	start := 1
	cnt := 9
	for n > cnt*l {
		n -= cnt * l
		l++
		cnt *= 10
		start *= 10
	}

	target := start + (n+l-1)/l - 1
	n -= (n - 1) / l * l
	x := target
	p := start
	for {
		n--
		if n == 0 {
			return x / p
		}

		x -= x / p * p
		p /= 10
	}
}
