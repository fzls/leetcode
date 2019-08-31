package leetcode

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	rStart := i

	for ; i >= 0 && s[i] != ' '; i-- {
	}
	rEnd := i

	return rStart - rEnd
}
