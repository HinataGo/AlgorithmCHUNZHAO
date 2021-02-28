package UnionFInd

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	var count int

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			// 发现一座岛屿，随后进行深搜
			if grid[x][y] == '1' {
				count++
				dfsGrid(x, y, grid)
			}
		}
	}
	return count
}

// 深搜思路， 以自己为中心 ，向左边搜 x-1，向右边搜 x+1 ,向上面搜y-1, 向下面搜 y+1 ，
// 搜索时第一考虑边界 因此 x >=0  x < len(grid)  y >= 0 y < len(grid[0])  因为是矩阵，所以四四方方
// 第二考虑 这个节点是否是 岛屿 1
// 如果发生越界， 或者非岛屿 也就是 0 ，直接 return
// 这里采用将访问后的岛屿修改成0 ，也就是 水， 但是如果实际要求不能修改grid ，那么还需要引入一个 外部变量 visited [][]bool  ,来给我们所有的节点进行访问标记
func dfsGrid(x, y int, grid [][]byte) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfsGrid(x-1, y, grid)
	dfsGrid(x+1, y, grid)
	dfsGrid(x, y-1, grid)
	dfsGrid(x, y+1, grid)
}
