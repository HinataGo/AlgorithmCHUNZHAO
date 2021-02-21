package thu

// 322. 零钱兑换
// 及时爬楼梯问题, 这样就很容易了
// 其他想法及时 广度优先遍历 ,第一个就是解, 最差的就是递归解法
func coinChange(coins []int, amount int) int {
	if amount == 0 || len(coins) == 0 {
		return 0
	}
	res := make([]int, amount+1)
	res[0] = 0
	for i := 1; i < amount+1; i++ {
		res[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				res[i] = min(res[i], res[i-coins[j]]+1)
			}
		}
	}
	if res[amount] > amount {
		return -1
	}
	return res[amount]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 爬楼梯
