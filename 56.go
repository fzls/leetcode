package leetcode

import "sort"

const (
	INTERVAL_BEGIN = 0
	INTERVAL_END   = 1
)

// 2019/10/30 23:37 by fzls
func _merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort by begin part
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][INTERVAL_BEGIN] < intervals[j][INTERVAL_BEGIN]
	})

	// now merge regions
	var mergedRegions [][]int

	n := len(intervals)
	intervalBegin := intervals[0][INTERVAL_BEGIN]
	intervalEnd := intervals[0][INTERVAL_END]
	for i := 1; i < n; i++ {
		interval := intervals[i]
		if interval[INTERVAL_BEGIN] <= intervalEnd {
			if interval[INTERVAL_END] > intervalEnd {
				intervalEnd = interval[INTERVAL_END]
			}
		} else {
			mergedRegions = append(mergedRegions, []int{intervalBegin, intervalEnd})
			intervalBegin = interval[INTERVAL_BEGIN]
			intervalEnd = interval[INTERVAL_END]
		}
	}
	mergedRegions = append(mergedRegions, []int{intervalBegin, intervalEnd})

	return mergedRegions
}
