package sat

import "math"

// 袖珍计算器法
// Exp返回ex，即x的底e指数。
// 特殊情况为：Exp（+ Inf）= + Inf Exp（NaN）= NaN非常大的值溢出到0或+ Inf。
// 很小的值下溢到1

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	res := int(math.Exp(0.5 * math.Log(float64(x))))
	if (res+1)*(res+1) <= x {
		return res + 1
	}
	return res
}
