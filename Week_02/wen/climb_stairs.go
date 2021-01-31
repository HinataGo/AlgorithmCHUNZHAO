package wen

// 同斐波那契额, 记住三个变量a b res 相互交换, 返回res
func climbStairs(n int) int {
	a, b, res := 0, 1, 0

	for i := 0; i < n; i++ {
		res = a + b
		a = b
		b = res
	}
	return res
}
