package leetcode

import "fmt"

const (
	MASK_OTHER = 0
	MASK_3     = 1
	MASK_5     = 1 << 1
)

func fizzBuzz(n int) []string {
	masks := make([]int, n+1)
	for i := 3; i <= n; i += 3 {
		masks[i] += MASK_3
	}
	for i := 5; i <= n; i += 5 {
		masks[i] += MASK_5
	}

	var res []string
	for i := 1; i <= n; i++ {
		if masks[i] == MASK_OTHER {
			res = append(res, fmt.Sprintf("%v", i))
		} else if masks[i] == MASK_3 {
			res = append(res, "Fizz")
		} else if masks[i] == MASK_5 {
			res = append(res, "Buzz")
		} else {
			res = append(res, "FizzBuzz")
		}
	}
	return res
}
