package leetcode

import "bytes"

var int2RomanMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var digits = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

// 2019/09/17 0:55 by fzls
func intToRoman(num int) string {
	if !(num >= 1 && num <= 3999) {
		return ""
	}

	var res bytes.Buffer

	for _, digit := range digits {
		for num >= digit {
			res.WriteString(int2RomanMap[digit])
			num -= digit
		}
	}

	return res.String()
}
