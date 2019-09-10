package leetcode

import "fmt"

func pow2(n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		t := pow2(n / 2)
		return t * t
	} else {
		return 2 * pow2(n-1)
	}
}

func _getPerms(sum int, max int, factor int, remainingOnes int, remainingDigits int, res *[]int) {
	if sum > max {
		return
	}
	if remainingOnes > remainingDigits {
		return
	}

	if remainingOnes == 0 {
		*res = append(*res, sum)
		return
	}

	// 这一位取1
	_getPerms(sum+factor, max, factor*2, remainingOnes-1, remainingDigits-1, res)
	// 这一位取0
	_getPerms(sum, max, factor*2, remainingOnes, remainingDigits-1, res)
}

// max: 和的最大值
// remainingOnes: 剩下的位中，还剩多少个1
// remainingDigits: 剩下的位数
func getPerms(max int, remainingOnes int, remainingDigits int) []int {
	if remainingDigits == 0 {
		return []int{0}
	}

	var res []int
	// 首位是0
	_getPerms(0, max, 1, remainingOnes, remainingDigits, &res)

	return res
}

func getTime(hour, min int) string {
	return fmt.Sprintf("%d:%02d", hour, min)
}

func readBinaryWatch(num int) []string {
	var times []string

	for nHour := 0; nHour <= 4; nHour++ {
		nMin := num - nHour
		if nMin < 0 || nMin > 6 {
			continue
		}

		// now hour has nHour led on(1), min has nMin led on(1)
		hours := getPerms(11, nHour, 4)
		mins := getPerms(59, nMin, 6)

		for _, hour := range hours {
			for _, min := range mins {
				times = append(times, getTime(hour, min))
			}
		}
	}

	return times
}
