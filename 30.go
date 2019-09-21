package leetcode

// 2019/09/21 21:42 by fzls
func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return nil
	}

	// 依次滑行s，看看是否匹配（优化可能需要KMP）
	var res []int

	l := len(words) * len(words[0])
	for i := 0; i+l-1 < len(s); i++ {
		if match(s[i:i+l], words) {
			res = append(res, i)
		}
	}

	return res
}

func match(s string, words []string) bool {
	used := make([]bool, len(words))
	return _match(s, words, used, len(words), len(words[0]), 0)
}

func _match(target string, words []string, used []bool, remain int, step int, matchedPart int) bool {
	if remain == 0 {
		return true
	}

	subTarget := target[matchedPart : matchedPart+step]
	for i := 0; i < len(words); i++ {
		if !used[i] && words[i] == subTarget {
			used[i] = true

			matched := _match(target, words, used, remain-1, step, matchedPart+step)
			if matched {
				return true
			}

			used[i] = false
		}
	}

	return false
}
