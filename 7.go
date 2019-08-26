package leetcode

func Run(){
	reverse(-1463847412)
}

const (
	INT32_MIN = -(1<<31)
	INT32_MIN_CHUSHU = INT32_MIN/10
	INT32_MIN_YUSHU = INT32_MIN%10
	INT32_MAX = (1<<31) - 1
	INT32_MAX_CHUSHU = INT32_MAX/10
	INT32_MAX_YUSHU = INT32_MAX%10
)

func isOverflow(n int, addon int) bool{
	if n > 0{
		return n > INT32_MAX_CHUSHU || n == INT32_MAX_CHUSHU && addon > INT32_MAX_YUSHU
	}else{
		return n < INT32_MIN_CHUSHU || n == INT32_MIN_CHUSHU && addon < INT32_MIN_YUSHU
	}
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		if isOverflow(res, x%10){
			return 0
		}
		res = 10*res + x%10
		x /= 10
	}

	return res
}
