package days

// 定义内部函数
// 内部函数赋值 给出具体操作
// 调用内部函数

// 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	var mid func(node *TreeNode)
	mid = func(node *TreeNode) {
		if node == nil {
			return
		}
		mid(node.Left)
		res = append(res, node.Val)
		mid(node.Right)
	}
	mid(root)
	return
}
