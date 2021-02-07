package thu

// 翻转二叉树 递归解决
// 递归解决很简单
// 先翻转左右子树
// 最后交换root 的两个子树,即完成递归翻转
// 切忌人肉递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
