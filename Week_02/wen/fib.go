package wen

// 配合爬楼楼梯,弄一个有趣的斐波那契
func fib(n int) int {
	a, res := 0, 1
	for n >= 0 {
		a, res = (a+res)%1000000007, a
		n--
	}
	return res
}
