package leetcode

// 其他解法：
// 分治法：LCP(s[i,j]) = LCP(LCP(s[i, m]), LCP(s[m+1, j]))
// 前缀树：Trie：根节点为无意义的一个节点
// 		构造：遍历每个字符串，当当前字符不在当前节点的子节点中，则新增一个该字符的子节点，否则设置当前节点为该子节点
// 		查询最大公共串：从根节点往下遍历，若子节点数目为1，则添加该节点的值到结果中，若不为1则返回当前结果

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var res []byte

	minLen := 0x7FFFFFFF
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

OUTER:
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[j-1][i] {
				break OUTER
			}
		}
		res = append(res, strs[0][i])
	}

	return string(res)
}
