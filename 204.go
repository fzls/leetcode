package leetcode

const (
	NOT_SURE  = 0
	IS_PRIME  = 1
	NOT_PRIME = 2
)

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}

	cnt := 0

	//使用素数筛
	isNotPrime := make([]bool, n) // 1~n-1
	for i := 2; i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		cnt++

		for j := 2 * i; j < n; j += i {
			isNotPrime[j] = true
		}
	}

	return cnt
}
