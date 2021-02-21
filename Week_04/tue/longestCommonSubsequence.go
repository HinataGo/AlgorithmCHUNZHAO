package tue

// 1143. 最长公共子序列
// 暴力思路一个一个遍历比较,
// 双重循环
// 动态规划思路,二维数组 res 记录
// 这里说一个有问题的思路, 就是不要想着下面从上面直接继承, 或者熊左边直接继承
// 主要这里有一个上学没弄清楚的误区
// for i ,j = 1,1
// if text1[i-1] == text2[j-1] --> res[i-1][j-1]+1
//
//   a b c d e
//  a 1 1 1 1 1
//  c 1 1 2 2 2
//  e 1 1 1 1 3

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	res := make2Slice(m+1, n+1)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = max(res[i-1][j], res[i][j-1])
			}
		}
	}
	return res[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func make2Slice(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
