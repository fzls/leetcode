package leetcode

import "sort"

// 2019/09/18 0:23 by fzls
func listIntToSet(nums []int) map[int]struct{} {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}
	return set
}

func listStringToSet(nums []string) map[string]struct{} {
	set := make(map[string]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}
	return set
}

func lessIntList(a, b []int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	} else if len(a) == 0 {
		return true
	} else if len(b) == 0 {
		return false
	} else {
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return lessIntList(a[1:], b[1:])
	}
}

func sortIntListList(data [][]int) {
	for i := 0; i < len(data); i++ {
		sort.Ints(data[i])
	}
	sort.Slice(data, func(i, j int) bool {
		return lessIntList(data[i], data[j])
	})
}

func sortIntPermutateList(data [][]int) {
	sort.Slice(data, func(i, j int) bool {
		return lessIntList(data[i], data[j])
	})
}

func lessStringList(a, b []string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	} else if len(a) == 0 {
		return true
	} else if len(b) == 0 {
		return false
	} else {
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return lessStringList(a[1:], b[1:])
	}
}

func sortStringList(data [][]string) {
	for i := 0; i < len(data); i++ {
		sort.Strings(data[i])
	}
	sort.Slice(data, func(i, j int) bool {
		return lessStringList(data[i], data[j])
	})
}

func sortStringPermutateList(data [][]string) {
	sort.Slice(data, func(i, j int) bool {
		return lessStringList(data[i], data[j])
	})
}
