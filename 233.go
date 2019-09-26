package leetcode

// 2019/09/27 0:54 by fzls
func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}

	res := 0

	pow10 := 1
	for x := n; x > 9; x /= 10 {
		pow10 *= 10
	}

	// left		digit	right
	// 12		3		45
	// 对于digit大于1的情况，这一位的1共有（12+1）*10^2个，分别对应left从0到12，right从0-99
	// 对于digit=1的情况，这一位的1共有12*10^2 + 45 + 1，分别对应（left从0到11，right从0-99）和left=12，right从0到45
	// 对于digit=0的情况，这一位的1共有12*10^2，分别对应left从0到11，right从0到45
	left := 0
	digit := 0
	right := n
	for ; pow10 > 0; pow10 /= 10 {
		digit = n / pow10 % 10

		right -= digit * pow10

		if digit == 0 {
			res += left * pow10
		} else if digit == 1 {
			res += left*pow10 + right + 1
		} else {
			res += (left + 1) * pow10
		}

		left = left*10 + digit
	}

	return res
}
