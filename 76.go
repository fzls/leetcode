package leetcode

import (
	"container/list"
)

// 2019/11/14 21:36 by fzls
type WindowItem struct {
	Char byte
	Pos  int
}

type Window struct {
	Begin int
	End   int
}

func minWindow(s string, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}

	// 用于统计还需要在窗口外匹配的字符集
	todos := make(map[byte]int)
	for _, c := range t {
		todos[byte(c)]++
	}

	// 最小窗口
	var min *Window
	// 当前的滑动窗口
	window := list.New()
	// 用于快速到找到当前元素在队列中的位置
	charElems := make(map[byte][]*list.Element)
	// 将新字符添加到窗口中
	addToWindow := func(pos int, char byte) {
		elem := window.PushBack(&WindowItem{
			Char: char,
			Pos:  pos,
		})
		charElems[char] = append(charElems[char], elem)
	}
	// 从窗口中移除一个元素
	popFromWindow := func(elem *list.Element) {
		window.Remove(elem)
		char := elem.Value.(*WindowItem).Char
		charElems[char] = charElems[char][1:]
	}
	// 参考题解先过滤一遍
	var filterItems []WindowItem
	for pos, c := range s {
		char := byte(c)
		if _, ok := todos[char]; ok {
			filterItems = append(filterItems, WindowItem{
				Char: char,
				Pos:  pos,
			})
		}
	}
	for _, item := range filterItems {
		pos, char := item.Pos, item.Char
		if _, ok := todos[char]; ok {
			// 如果在未匹配的T的字符集中有这个字符，则加入当前的匹配窗口列表中
			todos[char]--
			if todos[char] == 0 {
				delete(todos, char)
			}
			addToWindow(pos, char)
		} else {
			// 否则则将其从列表中移出一个，然后把这个放到队尾，从而使子串尽可能短
			popFromWindow(charElems[char][0])
			addToWindow(pos, char)
		}

		// 看看当前的窗口是否是一个符合要求的窗口
		if len(todos) == 0 {
			// 更新最小窗口
			newMindow := &Window{
				Begin: window.Front().Value.(*WindowItem).Pos,
				End:   window.Back().Value.(*WindowItem).Pos,
			}

			if min == nil || newMindow.End-newMindow.Begin < min.End-min.Begin {
				min = newMindow
			}

			// 删除窗口头部元素
			front := window.Front()
			popFromWindow(front)
			todos[front.Value.(*WindowItem).Char]++
		}
	}

	if min == nil {
		return ""
	}
	return s[min.Begin : min.End+1]
}
