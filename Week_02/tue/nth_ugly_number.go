package tue

// dp解法,取每次组合乘积的最小值
// dp取三者最小的一个树,记住是任意 2, 3,5 的公倍数,并且不重读的放进数组
func nthUglyNumber(n int) int {
	res := make([]int, 1)
	res[0] = 1
	ugly, i2, i3, i5 := 0, 0, 0, 0
	for i := 1; i <= n; i++ {
		ugly = min(min(res[i2]*2, res[i3]*3), res[i5]*5)
		res[i] = ugly
		if ugly == res[i2]*2 {
			i2++
		}
		if ugly == res[i3]*3 {
			i3++
		}
		if ugly == res[i5]*5 {
			i5++
		}
	}
	return res[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
