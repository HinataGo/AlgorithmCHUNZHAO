package thu

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
// 动态规划 - 改造
// 递归的时候多次求相同的值
// 自底向上
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
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
