package days

func minPathSum(grid [][]int) int {
	row, column := len(grid), len(grid[0])
	if row == 0 || column == 0 {
		return 0
	}
	res := make2Slice(row, column)
	res[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		res[i][0] = grid[i][0] + res[i-1][0]
	}
	for j := 1; j < column; j++ {
		res[0][j] = grid[0][j] + res[0][j-1]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			res[i][j] = min(res[i][j-1], res[i-1][j]) + grid[i][j]
		}
	}
	return res[row-1][column-1]
}

func make2Slice(row, column int) [][]int {
	res := make([][]int, row)
	for i := range res {
		res[i] = make([]int, column)
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
