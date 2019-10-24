package leetcode

import (
	"sort"
)

// 2019/10/19 0:50 by fzls
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	mp := make(map[string]int)

	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})

		ss := string(bs)
		idx, ok := mp[ss]
		if !ok {
			res = append(res, []string{})
			idx = len(res) - 1
			mp[ss] = idx
		}
		res[idx] = append(res[idx], s)
	}

	return res
}
