package days

// N叉树的后序遍历

func postorder(root *Node) (res []int) {
	if root == nil {
		return nil
	}
	for _, v := range root.Children {
		res = append(res, postorder(v)...)
	}
	res = append(res, root.Val)
	// return  res 这里 res 可加可不加
	return
}
