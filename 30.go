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
	wordsCntMap := make(map[string]int)
	for _, word := range words {
		wordsCntMap[word]++
	}
	for i := 0; i+l-1 < len(s); i++ {
		if match(s[i:i+l], wordsCntMap, lenWord) {
			res = append(res, i)
		}
	}

	return res
}

// 参考其他人的思路，将该段分成跟words一样的段落，然后比较双方各个单词数数目是否一致
func match(s string, wordsCntMap map[string]int, lenWord int) bool {
	cntMap := make(map[string]int)
	for i := 0; i < len(s); i += lenWord {
		cntMap[s[i:i+lenWord]]++
	}

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
