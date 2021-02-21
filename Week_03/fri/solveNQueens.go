package fri

import "strings"

// 51. N 皇后
// 每行每列 只能放一个皇后 同时保证 斜线上没有

// 与搜索全排列思路一致
// 需要剪枝
// 1. 可以先全排列,在判断 剪枝
// 2. 可以边判断合法,边保存排列结果
func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	cols := map[int]bool{}
	skim := map[int]bool{}
	press := map[int]bool{}

	res := make([][]string, 0)
	helper(0, bd, &res, n, &cols, &skim, &press)
	return res
}
func helper(r int, bd [][]string, res *[][]string, n int, cols, diag1, diag2 *map[int]bool) {
	if r == n {
		temp := make([]string, len(bd))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for c := 0; c < n; c++ {
		if !(*cols)[c] && !(*diag1)[r+c] && !(*diag2)[r-c] {
			bd[r][c] = "Q"
			(*cols)[c] = true
			(*diag1)[r+c] = true
			(*diag2)[r-c] = true
			helper(r+1, bd, res, n, cols, diag1, diag2)
			bd[r][c] = "."
			(*cols)[c] = false
			(*diag1)[r+c] = false
			(*diag2)[r-c] = false
		}
	}

}
