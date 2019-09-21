package leetcode

// 2019/09/21 21:42 by fzls
func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return nil
	}

	// 依次滑行s，看看是否匹配（优化可能需要KMP）
	var res []int

	lenWord := len(words[0])
	l := len(words) * lenWord

	// 参考其他人的思路，将该段分成跟words一样的段落，然后比较双方各个单词数数目是否一致
	wordsCntMap := make(map[string]int)
	for _, word := range words {
		wordsCntMap[word]++
	}
	for j := 0; j < lenWord; j++ {
		if j+l-1 >= len(s) {
			break
		}

		cntMap := make(map[string]int)
		for i := j; i < j+l; i += lenWord {
			cntMap[s[i:i+lenWord]]++
		}
		if match(cntMap, wordsCntMap) {
			res = append(res, j)
		}

		for i := j + l; i+lenWord-1 < len(s); i += lenWord {
			last := s[i-l : i-l+lenWord]
			next := s[i : i+lenWord]
			cntMap[last]--
			if cntMap[last] == 0 {
				delete(cntMap, last)
			}
			cntMap[next]++
			if match(cntMap, wordsCntMap) {
				res = append(res, i-l+lenWord)
			}
		}
	}

	return res
}

func match(cntMap, wordsCntMap map[string]int) bool {
	if len(cntMap) != len(wordsCntMap) {
		return false
	}
	for word, cnt := range wordsCntMap {
		if cntMap[word] != cnt {
			return false
		}
	}
	return true
}
