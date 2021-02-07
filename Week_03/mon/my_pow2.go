package mon

// 对求幂的次数进行分治, 随后得出结果
// 数据转换 n < 0时
// 再建 一个函数进行递归操作
// 首先判断 n== 0 则return 1.0
func myPow(x float64, n int) float64 {
	// 分治 的 奇偶性情况讨论
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return qPow(x, n)
}

func qPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	half := qPow(x, n/2)
	// 偶数情况
	if n%2 == 0 {
		return half * half
	} else {
		// 奇数情况
		return half * half * x
	}
}
