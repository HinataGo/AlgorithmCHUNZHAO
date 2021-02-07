package thu

// 二叉树 层序遍历 当模板用
// 递归版本
// 全局 ret 存结果
// 主要思想 1. 判断边界条件 特殊值
// 2. level == len(ret) 时 append 数据进 ret
// 3. ret[level] append 当前节点的val
// 4. 分开递归调用左右子树吗每次带上参数 level + 1

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res = make([][]int, 0)
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
