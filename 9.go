package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var arr []int
	for _x := x; _x > 0; _x /= 10 {
		arr = append(arr, _x%10)
	}

	i:=0
	j:=len(arr)-1
	for i<j {
		if arr[i] != arr[j]{
			return false
		}

		i++
		j--
	}

	return true
}
