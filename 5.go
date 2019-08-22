package leetcode

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	var (
		maxI = 0
		maxJ = 0
	)
	for i:=0; i < len(s); i++{
		for j:=i+1; j<len(s); j++{
			if j-i > maxJ-maxI && isPalindrome(s[i:j+1]) {
				maxI = i
				maxJ = j
			}
		}
	}

	return s[maxI:maxJ+1]
}

func isPalindrome(s string) bool {
	if len(s) <= 1{
		return true
	}

	i:=0
	j:=len(s) - 1
	for i<j {
		if s[i]!=s[j]{
			return false
		}

		i++
		j--
	}

	return true
}