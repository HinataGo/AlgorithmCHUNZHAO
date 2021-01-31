package mon

// 后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var post func(node *TreeNode)
	post = func(node *TreeNode) {
		if node == nil {
			return
		}
		post(node.Left)
		post(node.Right)
		res = append(res, node.Val)
	}
	post(root)
	return
}
