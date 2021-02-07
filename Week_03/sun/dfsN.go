package sun

func dfsN(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	// 下面两步骤顺序 目前是前序,互换为 后序遍历(没有中序)
	res = append(res, root.Val)
	// 特殊点 ,这里不要分开函数写了
	// 因为没有左右孩子区分,只有多个children ,因此这里,直接遍历左右孩子
	// 依次添加
	//
	for _, v := range root.Children {
		res = append(res, dfsN(v)...)
	}
	return res
}
