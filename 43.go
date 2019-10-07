package leetcode

// 2019/10/08 0:48 by fzls
func multiply(num1 string, num2 string) string {
	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}

	mul := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			sum := mul[i+j+1] + n1*n2
			mul[i+j+1] = sum % 10
			mul[i+j] += sum / 10
		}
	}

	var res []byte
	if mul[0] == 0 {
		res = make([]byte, 0, len(mul)-1)
	} else {
		res = make([]byte, 0, len(mul))
	}

	for i := 0; i < len(mul); i++ {
		if i == 0 && mul[i] == 0 {
			continue
		}

		res = append(res, byte(mul[i])+'0')
	}

	return string(res)
}
