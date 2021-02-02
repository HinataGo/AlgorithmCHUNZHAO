package thu

// 翻转二叉树 递归解决
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Right = left
	root.Left = right
	return root
}
