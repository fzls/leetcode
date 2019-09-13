package leetcode

func addStrings(num1 string, num2 string) string {
	var resStack []byte
	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0

	for i >= 0 && j >= 0 {
		sum := int(num1[i]-'0') + int(num2[j]-'0') + carry

		resStack = append(resStack, byte('0'+sum%10))
		carry = sum / 10

		i--
		j--
	}

	for i >= 0 {
		sum := int(num1[i]-'0') + carry

		resStack = append(resStack, byte('0'+sum%10))
		carry = sum / 10

		i--
	}
	for j >= 0 {
		sum := int(num2[j]-'0') + carry

		resStack = append(resStack, byte('0'+sum%10))
		carry = sum / 10

		j--
	}
	if carry > 0 {
		resStack = append(resStack, '1')
	}

	{
		i := 0
		j := len(resStack) - 1
		for i < j {
			resStack[i], resStack[j] = resStack[j], resStack[i]

			i++
			j--
		}
		return string(resStack)
	}
}
