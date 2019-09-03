package leetcode

func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = 26*res + int(s[i]-'A') + 1
	}
	return res
}
