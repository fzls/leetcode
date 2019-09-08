package leetcode

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	var temp byte
	for i < j {
		temp = s[i]
		s[i] = s[j]
		s[j] = temp

		i++
		j--
	}
}
