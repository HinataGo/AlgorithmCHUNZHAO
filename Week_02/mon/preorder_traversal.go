package mon

// 前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	var pre func(ndoe *TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		pre(node.Left)
		pre(node.Right)
	}
	pre(root)
	return
}
