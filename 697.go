package leetcode

// 2019/09/26 1:48 by fzls
type _count struct {
	cnt      int
	firstIdx int
	lastIdx  int
}

func (c *_count) Len() int {
	return c.lastIdx - c.firstIdx + 1
}

func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	counter := make(map[int]*_count)
	for idx, num := range nums {
		if cnt, ok := counter[num]; ok {
			cnt.cnt++
			cnt.lastIdx = idx
		} else {
			counter[num] = &_count{
				cnt:      1,
				firstIdx: idx,
				lastIdx:  idx,
			}
		}
	}

	var mc _count
	for _, cnt := range counter {
		if cnt.cnt > mc.cnt ||
			cnt.cnt == mc.cnt && cnt.Len() < mc.Len() {
			mc = *cnt
		}
	}

	return mc.Len()
}
