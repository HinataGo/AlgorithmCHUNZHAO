package wen

// 同斐波那契额, 记住三个变量a b res 相互交换, 返回res
// 优化写法,节约空间成本
// func climbStairs(n int) int {
// 	a, res := 0, 1
// 	for i := 0; i <= n; i++ {
// 		a, res = a+res, a
// 	}
// 	return a
// }
func climbStairs(n int) int {
	a, res := 0, 1
	for n >= 0 {
		a, res = a+res, a
		n--
	}
	return a
}
