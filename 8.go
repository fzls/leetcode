package leetcode

//请你来实现一个 atoi 函数，使其能将字符串转换成整数。
//
//首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
//
//当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；
// 假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
//
//该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
//
//注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
//
//在任何情况下，若函数不能进行有效的转换时，请返回 0。
//
//说明：
//
//假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/string-to-integer-atoi
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

const (
	MAX_INT32   = 1<<31 - 1
	MAX_INT32_L = MAX_INT32 / 10
	MAX_INT32_R = MAX_INT32 % 10

	MIN_INT32   = -(1 << 31)
	MIN_INT32_L = MIN_INT32 / 10
	MIN_INT32_R = MIN_INT32 % 10
)

func isSign(char byte) bool {
	return char == '+' || char == '-'
}

func toSign(sign byte) int {
	if sign == '+' {
		return 1
	} else {
		return -1
	}
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func toDigit(digit byte) int {
	return int(digit - '0')
}

func myAtoi(str string) int {
	res := 0
	sign := 1

	idx := 0
	for idx < len(str) && str[idx] == ' ' {
		idx++
	}

	// 现在idx是第一个非空字符的坐标
	// 如果全为空字符
	if idx == len(str) {
		return 0
	}

	// 如果第一个不是正负号或者数字
	if !isSign(str[idx]) && !isDigit(str[idx]) {
		return 0
	}

	if isSign(str[idx]) {
		sign = toSign(str[idx])
		idx++
	}
	for idx < len(str) && isDigit(str[idx]) {
		digit := toDigit(str[idx])

		// 判断是否溢出
		_res := sign * res
		if _res > MAX_INT32_L || _res == MAX_INT32_L && digit > MAX_INT32_R {
			return MAX_INT32
		}
		if _res < MIN_INT32_L || _res == MIN_INT32_L && -digit < MIN_INT32_R {
			return MIN_INT32
		}

		res = 10*res + digit

		idx++
	}

	return sign * res
}
