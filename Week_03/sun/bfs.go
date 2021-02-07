package sun

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var res [][]int

// bfs模板 二叉树
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
