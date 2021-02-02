package thu

// 二叉树的最大深度
// 先判断边界, 没有root 就是空 nil 也是0
// 有root 最后 +1 就是为了补上root那一层
// 深度 递归 ,取左右子树,最大的值
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
