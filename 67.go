package leetcode

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	addon := 0
	var resStack []byte
	for i >= 0 || j >= 0 {
		addLeft := 0
		if i >= 0 && a[i] == '1' {
			addLeft = 1
		}
		addRight := 0
		if j >= 0 && b[j] == '1' {
			addRight = 1
		}

		sum := addLeft + addRight + addon
		resStack = append(resStack, '0'+byte(sum&1))
		addon = sum >> 1

		i--
		j--
	}
	if addon != 0 {
		resStack = append(resStack, '1')
	}

	var res []byte
	for idx := len(resStack) - 1; idx >= 0; idx-- {
		res = append(res, resStack[idx])
	}

	return string(res)
}
