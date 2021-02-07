package days

// 扫雷游戏
// 点开一个点，要么是雷，要么不是：
//
//    点到 M，踩雷了，更新为X，游戏结束。
//    点到 E，空地，分两种情况：
//        周围 8 个格子有雷，更新为雷数。
//        周围 8 个格子没有雷，更新为 B，并继续探测这 8 个格子。
//
// 关键就是这个继续探测，就是一种搜索，有 DFS、BFS 两种实现方式
// DFS
//
// 无论是 DFS 还是 BFS，都要避免重复访问，这里怎么避免呢？
// 我们每访问一个 E，都会更新标记，不再是 E，下次再访问就直接返回。
// 不用再开辟空间去存访问过的节点，已经把访问信息就地存储了

var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		bfs(board, x, y)
	}
	return board
}

func bfs(board [][]byte, sx, sy int) {
	var queue [][]int
	vis := make([][]bool, len(board))
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, len(board[0]))
	}
	queue = append(queue, []int{sx, sy})
	vis[sx][sy] = true
	for i := 0; i < len(queue); i++ {
		cnt, x, y := 0, queue[i][0], queue[i][1]
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
				continue
			}
			// 不用判断 M，因为如果有 M 的话游戏已经结束了
			if board[tx][ty] == 'M' {
				cnt++
			}
		}
		if cnt > 0 {
			board[x][y] = byte(cnt + '0')
		} else {
			board[x][y] = 'B'
			for i := 0; i < 8; i++ {
				tx, ty := x+dirX[i], y+dirY[i]
				// 这里不需要在存在 B 的时候继续扩展，因为 B 之前被点击的时候已经被扩展过了
				if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' || vis[tx][ty] {
					continue
				}
				queue = append(queue, []int{tx, ty})
				vis[tx][ty] = true
			}
		}
	}
}
