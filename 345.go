package leetcode

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func reverseVowels(s string) string {
	bytes := []byte(s)

	i := 0
	j := len(s) - 1
	var temp byte
	for i < j {
		for i < j && !isVowel(bytes[i]) {
			i++
		}
		for i < j && !isVowel(bytes[j]) {
			j--
		}
		if i < j {
			temp = bytes[i]
			bytes[i] = bytes[j]
			bytes[j] = temp

			i++
			j--
		}
	}

	return string(bytes)
}
