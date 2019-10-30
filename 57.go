package leetcode

import "sort"

// 2019/10/30 23:58 by fzls
func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)

	// 将新区间按照begin的大小插入到合适的地方，保证所有其左侧的begin都小于等于它，右侧均大于它
	insertIdx := sort.Search(n, func(i int) bool {
		return intervals[i][0] > newInterval[0]
	})

	// 结果集
	var res [][]int

	// 根据定义，在插入位置之前的必定是属于最终结果集的
	res = append(res, intervals[:insertIdx]...)

	// 处理新插入的区间与前一个区间是否需要合并
	if len(res) != 0 && newInterval[0] <= res[len(res)-1][1] {
		if newInterval[1] > res[len(res)-1][1] {
			res[len(res)-1][1] = newInterval[1]
		}
	} else {
		res = append(res, newInterval)
	}

	// 处理新插入的区间是否需要与后续的区间进行合并
	for idx := insertIdx; idx < n; idx++ {
		if intervals[idx][0] <= res[len(res)-1][1] {
			// 该部分需要与新插入的合并到一起
			if intervals[idx][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[idx][1]
			}
		} else {
			// 新插入的区间已合并完成，后续均已保证不会重叠，可直接返回
			res = append(res, intervals[idx:]...)
			break
		}
	}

	return res
}
