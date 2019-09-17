package leetcode

var digit2Alphas = []string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

// 2019/09/18 0:21 by fzls
func letterCombinations(digits string) []string {
	return _letterCombinations(digits, []string{})
}

func _letterCombinations(digits string, preCombinations []string) []string {
	if len(digits) == 0 {
		return preCombinations
	}

	if len(preCombinations) == 0 {
		preCombinations = append(preCombinations, "")
	}

	var combinations []string
	for _, pc := range preCombinations {
		alphas := digit2Alphas[digits[0]-'0']
		for i := 0; i < len(alphas); i++ {
			combinations = append(combinations, pc+alphas[i:i+1])
		}
	}

	return _letterCombinations(digits[1:], combinations)
}
