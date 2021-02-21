package mon

/*func minimumTotal(triangle [][]int) int {
    return find(triangle,0,0)
}

func find(arr [][]int, i,j int)int{
    if len(arr) == 0{
        return 0
    }
    return min(find(arr,i,j), find(arr,i+1,j+1)) + arr[i][j]
}
func min(a,b int) int{
    if a > b{
        return b
    }
    return a
}
*/
// 递归--- out of memory  stack炸了
// 递归记忆化
// 动态规划 - 改造
// 动态规划空间优化
// 递归的时候多次求相同的值
/*
定义二维 dp 数组，将解法二中「自顶向下的递归」改为「自底向上的递推」。
1、状态定义：
dp[i][j] 表示从点 (i,j) 到底边的最小路径和。
2、状态转移：
dp[i][j] = min( dp[i+1][j], dp[i+1][j+1] ) + triangle[i][j]

*/
// func minimumTotal(triangle [][]int) int {
// 	n := len(triangle)
// 	if n == 0 {
// 		return 0
// 	}
// 	res := make([][]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		res[i] = make([]int, n+1)
// 	}
// 	for i := n - 1; i >= 0; i-- {
// 		for j := 0; j <= i; j++ {
// 			res[i][j] = min(res[i+1][j], res[i+1][j+1]) + triangle[i][j]
// 		}
// 	}
// 	return res[0][0]
// }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 递归优化
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	res := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			res[j] = min(res[j], res[j+1]) + triangle[i][j]
		}
	}
	return res[0]
}
