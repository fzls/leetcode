package leetcode

// 2019/09/13 20:48 by fzls
func countSegments(s string) int {
	cnt := 0

	i := 0
	for i < len(s) {
		// 找到第一个不是空格的字符，并把cnt+1
		for ; i < len(s) && s[i] == ' '; i++ {

		}
		if i != len(s) {
			cnt++
		}

		// 找到之后的第一个空格，表示单词末尾
		for ; i < len(s) && s[i] != ' '; i++ {

		}
	}

	return cnt
}
