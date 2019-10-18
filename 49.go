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
		if idx, ok := mp[ss]; ok {
			res[idx] = append(res[idx], s)
		} else {
			res = append(res, []string{s})
			mp[ss] = len(res) - 1
		}
	}

	return res
}
