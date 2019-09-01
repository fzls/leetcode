package leetcode

var mapRomanToInt map[byte]int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res := 0
	chars := []byte(s)
	for idx, c := range chars {
		val := mapRomanToInt[c]
		if idx+1 < len(chars) {
			nextVal := chars[idx+1]

			if c == 'I' && (nextVal == 'V' || nextVal == 'X') ||
				c == 'X' && (nextVal == 'L' || nextVal == 'C') ||
				c == 'C' && (nextVal == 'D' || nextVal == 'M') {
				val *= -1
			}
		}

		res += val
	}

	return res
}
