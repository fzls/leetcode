package leetcode

import "sort"

// 2019/10/30 23:58 by fzls
func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)

	// 将新区间按照begin的大小插入到合适的地方，保证所有其左侧的begin都小于等于它，右侧均大于它
	insertIdx := sort.Search(n, func(i int) bool {
		return intervals[i][0] > newInterval[0]
	})

	var res [][]int

	for idx := 0; idx < insertIdx; idx++ {
		res = append(res, intervals[idx])
	}

	if len(res) != 0 && newInterval[0] <= res[len(res)-1][1] {
		if newInterval[1] > res[len(res)-1][1] {
			res[len(res)-1][1] = newInterval[1]
		}
	} else {
		res = append(res, newInterval)
	}

	for idx := insertIdx; idx < n; idx++ {
		if intervals[idx][0] <= res[len(res)-1][1] {
			if intervals[idx][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[idx][1]
			}
		} else {
			res = append(res, intervals[idx])
		}
	}

	return res
}
