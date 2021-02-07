package tue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

// bfs
func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	bfs(root, 0)
	return res
}
func bfs(root *TreeNode, level int) [][]int {
	if root == nil {
		return res
	}
	if level == len(res) {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}
	res = bfs(root.Left, level+1)
	res = bfs(root.Right, level+1)
	return res
}
