package sun

type Node struct {
	Val      int
	Children []*Node
}

func levelOrderN(root *Node) [][]int {
	res = make([][]int, 0)
	bfsN(root, 0)
	return res
}
func bfsN(root *Node, level int) {
	if root == nil {
		return
	}
	if level == len(res) {
		res = append(res, []int{})
	}
	// 与二叉树的层序遍历不同之处在于这里
	// 此时,我们发现,还需要对孩子结点单独遍历,斌企鹅使用bfs
	res[level] = append(res[level], root.Val)
	for _, v := range root.Children {
		bfsN(v, level+1)
	}
}
