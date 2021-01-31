package wen

// 括号生成
// 简单来说，在求N个括号的排列组合时，把第N种情况（也就是N个括号排列组合）
// 视为单独拿一个括号E出来，剩下的N-1个括号分为两部分，
// P个括号和Q个括号，P+Q=N-1，然后这两部分分别处于括号E内和括号E的右边，
// 各自进行括号的排列组合。由于我们是一步步计算得到N个括号的情况的，
// 所以小于等于N-1个括号的排列组合方式我们是已知的（用合适的数据结构存储，
// 方便后续调用，且在存储时可利用特定数据结构实现题目某些要求，如排序，去重等），
// 且P+Q=N-1，P和Q是小于等于N-1的，所以我们能直接得到P个和Q个括号的情况，
// 进而得到N个括号的结果！
//
// 这个算法主要的基点就是将排列组合的情况分为了括号内和括号外这两种情况， 且仅存在两种情况！
// 至于为什么，原因在于楼主的算法的前提是单独拿出来的括号E的左边在N个括号所有排列组合情况中都是处于最左边，
// 所以不存在括号位于括号E的左边的情况。
// 因此，N-1个括号（拿出了括号E）仅可能分布于括号E内和括号E外，
// 分为两种子情况讨论！ 这种思想还可以应用于其他类似的题的求解中，
// 即怎样合理高效的利用前面步骤的计算结果得出当前步骤结果，从而得出最终结果
func generateParenthesis(n int) []string {
	res := []string{}

	var dfs func(lRemain int, rRemain int, path string)
	dfs = func(lRemain int, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
			return
		}
		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		if lRemain < rRemain {
			dfs(lRemain, rRemain-1, path+")")
		}
	}

	dfs(n, n, "")
	return res
}
