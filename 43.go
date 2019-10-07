package leetcode

import (
	"fmt"
)

// 2019/10/08 0:48 by fzls
func multiply(num1 string, num2 string) string {
	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}

	var adds []string

	for i := 0; i < len(num2); i++ {
		// 计算num2[i] * num1
		var mul []byte
		var carry, prod, digit byte
		for j := len(num1) - 1; j >= 0; j-- {
			prod = (num2[i]-'0')*(num1[j]-'0') + carry

			carry = prod / 10
			digit = prod % 10
			mul = append([]byte{digit + '0'}, mul...)
		}
		if carry > 0 {
			mul = append([]byte{carry + '0'}, mul...)
		}
		for z := 0; z < len(num2)-i-1; z++ {
			mul = append(mul, '0')
		}
		adds = append(adds, string(mul))
	}

	fmt.Println(adds)

	return addStr(adds)
}

func addStr(toAdds []string) string {
	if len(toAdds) == 0 {
		return "0"
	}
	if len(toAdds) == 1 {
		return toAdds[0]
	}

	sum := toAdds[0]
	for i := 1; i < len(toAdds); i++ {
		sum = addStr2(sum, toAdds[i])
	}
	return sum
}

func addStr2(num1 string, num2 string) string {
	var sum []byte

	i := len(num1) - 1
	j := len(num2) - 1
	var carry, digit byte
	for i >= 0 || j >= 0 {
		s := carry
		if i >= 0 {
			s += num1[i] - '0'
		}
		if j >= 0 {
			s += num2[j] - '0'
		}

		carry = s / 10
		digit = s % 10
		sum = append([]byte{digit + '0'}, sum...)

		i--
		j--
	}
	if carry != 0 {
		sum = append([]byte{'1'}, sum...)
	}
	return string(sum)
}
