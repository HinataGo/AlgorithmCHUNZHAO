package sat

// 二分搜索
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 1, x/2
	for l < r {
		mid := (l + r + 1) >> 1
		square := mid * mid
		if square > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return int(l)
}
