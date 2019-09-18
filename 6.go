package leetcode

//func Run() {
//	fmt.Println(convert("A", 1))
//	//"LCIRETOESIIGEDHN"
//}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := make([]byte, len(s))
	idx := 0

	m := len(s)
	// z字形首行，间隔2(n-1)
	for j := 1; j <= m; j += 2 * (numRows - 1) {
		res[idx] = s[j-1]
		idx++
	}

	// z字形中间行，第i行，则间距依次为2*(n-i), 2*(i-1)
	for i := 2; i < numRows; i++ {
		j := i
		for ; ; {
			if j <= m {
				res[idx] = s[j-1]
				idx++
			} else {
				break
			}
			j += 2 * (numRows - i)
			if j <= m {
				res[idx] = s[j-1]
				idx++
			} else {
				break
			}
			j += 2 * (i - 1)
		}
	}

	// z字形最后一行，间距为2*(n-1)
	for j := numRows; j <= m; j += 2 * (numRows - 1) {
		res[idx] = s[j-1]
		idx++
	}

	return string(res)
}
