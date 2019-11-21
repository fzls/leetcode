package leetcode

import (
	"sort"
)

// 2019/11/21 22:12 by fzls
func largestRectangleArea(heights []int) int {
	return _largestRectangleArea(heights)
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	indexByHeights := make([]int, len(heights))
	for i := 0; i < len(indexByHeights); i++ {
		indexByHeights[i] = i
	}
	sort.Slice(indexByHeights, func(i, j int) bool {
		return heights[indexByHeights[i]] < heights[indexByHeights[j]]
	})

	maxArea := 0
	regions := [][]int{
		{0, len(heights) - 1},
	}
	for _, idx := range indexByHeights {
		regionIdx := bSearch(regions, idx)
		region := regions[regionIdx]
		area := (region[1] - region[0] + 1) * heights[idx]
		if area > maxArea {
			maxArea = area
		}
		newRegions := append([][]int{}, regions[:regionIdx]...)
		if region[0] != idx {
			newRegions = append(newRegions, []int{region[0], idx - 1})
		}
		if region[1] != idx {
			newRegions = append(newRegions, []int{idx + 1, region[1]})
		}
		if regionIdx+1 < len(regions) {
			newRegions = append(newRegions, regions[regionIdx+1:]...)
		}
		regions = newRegions
	}

	return maxArea
}

func bSearch(regions [][]int, idx int) int {
	low, high := 0, len(regions)-1
	for low <= high {
		mid := (low + high) / 2
		region := regions[mid]
		if region[0] <= idx && idx <= region[1] {
			return mid
		} else if region[1] < idx {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// 单调栈
func _largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := []int{-1}
	for idx, height := range heights {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= height {
			area := heights[stack[len(stack)-1]] * (idx - stack[len(stack)-2] - 1)
			if area > maxArea {
				maxArea = area
			}
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, idx)
	}
	for stack[len(stack)-1] != -1 {
		area := heights[stack[len(stack)-1]] * (len(heights) - stack[len(stack)-2] - 1)
		if area > maxArea {
			maxArea = area
		}
		stack = stack[:len(stack)-1]
	}
	return maxArea
}
