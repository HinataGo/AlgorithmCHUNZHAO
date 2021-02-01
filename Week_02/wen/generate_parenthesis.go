package wen

// 括号生成
// 剪枝算法
// 定义全局变量
var res []string

func generateParenthesis(n int) []string {
	// 每次使用前都得初始化,否则之前的数据残留会产生错误
	res = make([]string, 0)
	generate(0, 0, n, "")
	return res
}
func generate(left, right, n int, s string) {
	// terminator 终止条件
	if left == n && right == n {
		res = append(res, s)
		return
	}
	// process 过程
	// drill down
	// 添加左括号
	if left < n {
		generate(left+1, right, n, s+"(")
	}
	// 添加右括号
	if left > right && right < n {
		generate(left, right+1, n, s+")")
	}
	// reverse states
}
