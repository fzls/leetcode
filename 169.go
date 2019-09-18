package leetcode

func majorityElement(nums []int) int {
	// 参考之前在编程之美中看到的每次干掉一对不同数的思路
	var stack []int

	for _, num := range nums {
		if len(stack) == 0 || stack[len(stack)-1] == num {
			stack = append(stack, num)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return stack[0]
}
