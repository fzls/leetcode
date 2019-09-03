package leetcode

func convertToTitle(n int) string {
	var stack []byte
	for n != 0 {
		stack = append(stack, byte((n-1)%26)+'A')
		n = (n - 1) / 26
	}

	// 翻转stack，就是结果了
	res := make([]byte, len(stack))
	for i := 0; i < len(res); i++ {
		res[i] = stack[len(stack)-1-i]
	}

	return string(res)
}
