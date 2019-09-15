package leetcode

type __key struct {
	s, p string
}

var cache = make(map[__key]bool)

// 2019/09/14 17:26 by fzls
func isMatch(s string, p string) (res bool) {
	// 对重复计算进行缓存
	if matched, ok := cache[__key{s, p}]; ok {
		return matched
	}

	defer func() {
		cache[__key{s, p}] = res
	}()

	// 如果pattern已全部消耗，则只需判断字符串是否也已消耗完
	if len(p) == 0 {
		return len(s) == 0
	}

	// 判断第一个字符是否匹配
	firstMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		// 判断有*通配符的情况
		return isMatch(s, p[2:]) || // 不匹配*
			(firstMatch && isMatch(s[1:], p)) // 消耗掉一个源字符串的字符
	} else {
		// 无通配符
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

// 参考题解使用状态机来实现正则表达式
func isMatchWithStateMachine(s string, p string) bool {
	init := &State{
		Val: '>',
	}
	init.Size = genStates(init, p, 0)
	return check(init, s, 0)
}

type State struct {
	Val      byte
	Parent   *State
	Children map[byte][]*State
	End      bool
	Size     int
}

func (s *State) Append(val byte, child *State) {
	if s.Children == nil {
		s.Children = make(map[byte][]*State)
	}
	for _, v := range s.Children[val] {
		if v == child {
			return
		}
	}
	s.Children[val] = append(s.Children[val], child)
}

func genStates(now *State, str string, idx int) int {
	if idx >= len(str) {
		now.End = true
		return now.Size // todo：这个返回值其实还不太理解
	}

	// 状态转移
	var nowState *State
	switch str[idx] {
	case '*':
		now.Size = 0
		now.Append(now.Val, now)
		nowState = now
	default:
		state := &State{}
		state.Val = str[idx]
		state.Size = 1
		state.Parent = now
		now.Append(str[idx], state)
		nowState = state
	}

	ret := genStates(nowState, str, idx+1)
	if ret == 0 {
		now.End = true
	}

	// 构建回环
	addParent := now // 从当前传入的节点开始
	for addParent.Parent != nil && addParent.Size == 0 {
		// 如果其父节点不为空，且该节点大小为0，则其父节点可以直接连接到本节点，也就是添加一个child
		addParent.Parent.Append(nowState.Val, nowState)
		// 依次向上迭代
		addParent = addParent.Parent
	}

	return now.Size + ret
}

func check(now *State, str string, idx int) bool {
	if idx >= len(str) {
		return now.End
	}

	// 下一步可走的包括.的子状态和当前字符的子状态
	states := append(now.Children['.'], now.Children[str[idx]]...)
	for _, state := range states {
		if check(state, str, idx+1) {
			return true
		}
	}

	// 所有状态转移均走不到终点，说明不能匹配
	return false
}
