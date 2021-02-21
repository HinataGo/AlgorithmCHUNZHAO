package tue

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
