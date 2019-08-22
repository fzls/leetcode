package leetcode

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var (
		maxI = 0
		maxJ = 0
	)
	var isPalindrome [1000][1000]bool
	// f(i,i)
	for i := 0; i < len(s); i++ {
		isPalindrome[i][i] = true
	}
	// f(i,i+1)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			isPalindrome[i][i+1] = true
			if maxJ-maxI < 1 {
				maxI = i
				maxJ = i + 1
			}
		} else {
			isPalindrome[i][i+1] = false
		}
	}

	// f(i,j) = false if s[i]!=s[j], f(i+1,j-1) if s[i] == s[j], for any 0<=i<n, i+2<=j<n
	for step := 2; step < len(s); step++ {
		i := 0
		j := i + step
		for j < len(s) {
			if s[i] == s[j] {
				isPalindrome[i][j] = isPalindrome[i+1][j-1]
				if maxJ-maxI < step && isPalindrome[i][j] {
					maxI = i
					maxJ = j
				}
			} else {
				isPalindrome[i][j] = false
			}

			i++
			j++
		}
	}

	return s[maxI : maxJ+1]
}
