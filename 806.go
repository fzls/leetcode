package leetcode

// 2019/10/18 0:52 by fzls
func numberOfLines(widths []int, S string) []int {
	if len(widths) != 26 {
		return nil
	}

	lines := 1
	remainingWidth := 100
	for _, char := range S {
		w := widths[char-'a']
		if remainingWidth < w {
			remainingWidth = 100
			lines++
		}

		remainingWidth -= w
	}

	return []int{lines, 100 - remainingWidth}
}
