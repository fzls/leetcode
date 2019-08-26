package leetcode

func Run(){
	reverse(1534236469)
}

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	res := 0
	for x > 0 {
		res = 10*res + x%10
		x /= 10
	}

	res*=sign

	if res < -(1<<31) || res > (1<<31) - 1 {
		return 0
	}

	return res
}
