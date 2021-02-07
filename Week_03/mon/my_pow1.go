package mon

// 其实go 底层有math.Pow()
// 判断条件
//  利用分治 折半 ,指数递增的思想

// n 为奇数 时 多补充一个x值乘积
// 偶次乘积, x 的2 4 8 16 次方依次递增,而同时 i 的值也在不断折半, 知道等于0
// 这里n 是整数,编程里对整数 带有小数时是只截取整数,不保留小数
// 随后判断, 如果 n < 0  res求倒数,否则 直接返回 res

func myPow2(x float64, n int) float64 {
	var res = 1.0
	for i := n; i != 0; i /= 2 {

		if i%2 != 0 {
			res *= x
		}
		x *= x
	}
	if n < 0 {
		return 1.0 / res
	}
	return res
}
