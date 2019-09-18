package leetcode

var mapRightToLeft = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	if len(s)&1 == 1 {
		return false
	}

	var stack []byte
	for _, c := range []byte(s) {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack)==0 || stack[len(stack)-1] != mapRightToLeft[c] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
