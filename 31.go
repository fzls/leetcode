package leetcode

// 2019/09/22 23:10 by fzls
// 除了举例，别忘了画图，本题单单举例可能看不出太多东西，但是一画图就很清晰了
// 最小的排列是完全顺序的，最大的排列是完全逆序的，也就是说排列递增的过程其实是让逆序对增加的过程，所以可以考虑转化为如果使当前排列的逆序对增加最小值
// 参考题解得到的理解：对于一个完全逆序的排列，我们无法找到比其更大的排列；所以我们从右往左找到第一个不是逆序的位置i,i+1（假设有）
// 则 a[0],... a[i], a[i+1], ... a[n-1], 其中a[i]<a[i+1], a[i+2] >= a[i+3], ... a[n-2] >= a[n-1]
// 然后我们从右往左，找到在i右侧第一个比i大的数，并将其与i交换位置，然后将i后方的全部翻转，则可以得到下一个排列
// 画一个示意图可以很容易理解上面的做法
func nextPermutation(nums []int) {
	last := len(nums) - 1
	i := last - 1
	for ; i >= 0; i-- {
		// 找到第一个连续顺序队 i, i+1
		if nums[i] < nums[i+1] {
			// 二分查找最后一个与该顺序对左元素构成顺序对的元素
			j := bs(i+1, last, func(idx int) bool {
				return nums[idx] > nums[i]
			})
			// 交换这组逆序对
			nums[j], nums[i] = nums[i], nums[j]
			// 将剩余部分排序，使该部分逆序数最小
			break
		}
	}

	// 完全逆序
	i++
	j := last
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// f(low-1) = true, f(high+1) = false
// find j such that f(i) = true for all low-1 < i <= j, f(i)=false for all j< i < high+1
func bs(low int, high int, f func(i int) bool) int {
	for low <= high {
		j := int(uint(low+high) >> 1)
		if f(j) {
			low = j + 1
			// for i < low, we have f(i) == true
		} else {
			high = j - 1
			// for i > high, we have f(i) == false
		}
	}
	return low - 1
}
