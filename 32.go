package leetcode

// 2019/09/28 1:18 by fzls
type Ele struct {
	CurLen       int
	UnpairedLeft int
}

func longestValidParentheses(s string) int {
	maxLen := 0

	var stack []byte
	var eleStack []Ele = []Ele{{0, 0}}
	for _, char := range []byte(s) {
		lastEle := eleStack[len(eleStack)-1]
		if char == '(' {
			// (
			stack = append(stack, char)
			eleStack = append(eleStack, Ele{CurLen: 0, UnpairedLeft: lastEle.UnpairedLeft + 1})
		} else {
			// )
			if len(stack) == 0 {
				eleStack = []Ele{{0, 0}}
				continue
			}

			stack = stack[:len(stack)-1]
			last2Ele := eleStack[len(eleStack)-2]
			eleStack = eleStack[:len(eleStack)-2]
			ele := Ele{CurLen: last2Ele.CurLen + lastEle.CurLen + 2, UnpairedLeft: lastEle.UnpairedLeft - 1}
			eleStack = append(eleStack, ele)

			if ele.CurLen > maxLen {
				maxLen = ele.CurLen
			}
		}
	}

	return maxLen
}
