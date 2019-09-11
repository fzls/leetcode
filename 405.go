package leetcode

var digit2Hex = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f',
}

func positiveHex(num uint32) string {
	var res []byte
	for num > 0 {
		res = append([]byte{digit2Hex[num%16]}, res...)
		num /= 16
	}
	return string(res)
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num > 0 {
		return positiveHex(uint32(num))
	}
	// negative
	x := uint32(^uint32(-num)) + 1
	return positiveHex(x)
}
