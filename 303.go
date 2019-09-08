package leetcode

type NumArray struct {
	nums []int

	kSums []int
}

func Constructor___(nums []int) NumArray {
	na := NumArray{
		nums: nums,
	}

	na.preCompute()

	return na
}

func (this *NumArray) preCompute() {
	sum := 0
	for i := 0; i < len(this.nums); i++ {
		sum += this.nums[i]

		this.kSums = append(this.kSums, sum)
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.kSums[j] - this.kSums[i] + this.nums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
