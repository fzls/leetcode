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

func _getPerms(sum int, max int, lastDigit int, remainNumberOfDigit int, res *[]int) {
	if sum > max {
		return
	}

	if remainNumberOfDigit == 0 {
		*res = append(*res, sum)
		return
	}

	// 这一位1至少是上一位左移一位
	t := (pow2(remainNumberOfDigit) - 1)
	for thisDigit := lastDigit * 2; t*thisDigit <= max; thisDigit *= 2 {
		_getPerms(sum+thisDigit, max, thisDigit, remainNumberOfDigit-1, res)
	}
}

func getPerms(max int, remainNumberOfDigit int) []int {
	if remainNumberOfDigit == 0 {
		return []int{0}
	}

	var res []int
	// 首位是1
	_getPerms(1, max, 1, remainNumberOfDigit-1, &res)
	// 首位是0
	_getPerms(0, max, 1, remainNumberOfDigit, &res)

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
		hours := getPerms(11, nHour)
		mins := getPerms(59, nMin)

		for _, hour := range hours {
			for _, min := range mins {
				times = append(times, getTime(hour, min))
			}
		}
	}

	return times
}
