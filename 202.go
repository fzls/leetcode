package leetcode

func isHappy(n int) bool {
	set := make(map[int]struct{})

	for n != 1 {
		if _, ok := set[n]; ok {
			return false
		}
		set[n] = struct{}{}
		res := 0
		for n != 0 {
			res += (n % 10) * (n % 10)
			n /= 10
		}
		n = res
	}

	return true
}
