package sat

// 使用牛顿迭代法求平方根
// (x+a/x)/2
// func mySqrt3(x int) int {
// 	if x == 0 {
// 		return 0
// 	}
// 	C, x0 := float64(x), float64(x)
// 	for {
// 		xi := 0.5 * (x0 + C/x0)
// 		if math.Abs(x0-xi) < 1e-7 {
// 			break
// 		}
// 		x0 = xi
// 	}
// 	return int(x0)
// }

func mySqrt3(x int) int {
	var a int = x
	for a*a > x {
		a = (a + x/a) / 2
	}
	return a
}
