package leetcode

// 2019/09/24 3:03 by fzls
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	negative := false
	if num < 0 {
		negative = true
		num = -num
	}

	var res []byte
	for num > 0 {
		res = append([]byte{'0' + byte(num%7)}, res...)

		num /= 7
	}
	if negative {
		res = append([]byte{'-'}, res...)
	}
	return string(res)
}
