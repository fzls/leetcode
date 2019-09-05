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
	isPrimes := make([]int, n) // 1~n-1
	isPrimes[1] = NOT_PRIME
	for i := 2; i < n; i++ {
		if isPrimes[i] == NOT_SURE {
			isPrimes[i] = IS_PRIME
			cnt++
		}

		for j := 2 * i; j < n; j += i {
			if isPrimes[j] == NOT_SURE {
				isPrimes[j] = NOT_PRIME
			}
		}
	}

	return cnt
}
