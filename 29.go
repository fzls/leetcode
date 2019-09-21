package leetcode

// 2019/09/21 21:06 by fzls
func divide(dividend int, divisor int) int {
	if dividend == -1<<31 && divisor == -1 {
		return (1 << 31) - 1
	}

	negative := false
	dd := dividend
	if dividend < 0 {
		dd = ^dividend + 1
		negative = !negative
	}
	dr := divisor
	if divisor < 0 {
		dr = ^divisor + 1
		negative = !negative
	}

	res := _divide(dd, dr)
	if negative {
		res = ^res + 1
	}
	return res
}

func _divide(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}

	// dividend >= 0, divisor > 0
	res := 0

	// 假设dividend = (fi*2^i + ... fo*2^0)divisor，计算出最大的这个fi，然后每次减去2^fi
	factor := uint(0)
	for (dividend >> factor) >= divisor {
		factor++
	}
	factor--

	for {
		if dividend>>factor >= divisor {
			dividend -= divisor << factor
			res += 1 << factor
		}
		if factor == 0 {
			break
		}
		factor--
	}
	return res
}
