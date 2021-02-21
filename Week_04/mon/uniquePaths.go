package mon

// 62. 不同路径

// 分析题目,找子重复子问题,建立状态转移方程 确定记忆化数据结构记忆化res
//
// res[0][0] = 1
// res[i][j] = res[i-1][j] + res[i][j-1] + 1
// res[0][j] = 1
// res[i][0] = 1

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	res := make2Slice(m, n)
	res[0][0] = 1
	for i := 1; i < m; i++ {
		res[i][0] = 1
	}
	for j := 1; j < n; j++ {
		res[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]
}

// 二维数组申明方式 res := [m][n]int{}, 这里m,n不能是变量.必须是常量,或者干脆是数字
func make2Slice(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
