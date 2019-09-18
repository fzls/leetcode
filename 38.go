package leetcode

var cntCache = make(map[int][]byte)

func cntToBytes(cnt int) []byte {
	if cache := cntCache[cnt]; cache != nil{
		return cache
	}

	var res []byte

	var stack []byte
	if cnt == 0 {
		stack = append(stack, 0)
	} else {
		for cnt > 0 {
			stack = append(stack, '0'+byte(cnt%10))
			cnt /= 10
		}
	}
	for idx := len(stack) - 1; idx >= 0; idx-- {
		res = append(res, stack[idx])
	}

	cntCache[cnt] = res

	return res
}

func countAndSay(n int) string {
	res := "1"
	if n == 1 {
		return res
	}

	for r := 2; r <= n; r++ {
		// 根据r-1个报数计算第r个报数
		var next []byte
		i := 0

		for j := i; j < len(res); j++ {
			if res[j] != res[i] {
				// append cnt
				next = append(next, cntToBytes(j-i)...)
				// append val
				next = append(next, res[i])

				// update i
				i = j
			}
		}

		// 记录最后一段相等的数字
		next = append(next, cntToBytes(len(res)-i)...)
		next = append(next, res[i])

		res = string(next)
	}

	return res
}
