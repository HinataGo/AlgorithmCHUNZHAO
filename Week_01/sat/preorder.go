package sat

// N叉树遍历 前序

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	res = append(res, root.Val)
	for _, v := range root.Children {
		res = append(res, preorder(v)...)
	}
	return res
}

// func preorder(root *Node) (res []int ){
// 	var pre func(ndoe * Node)
// 	pre = func(node *Node){
// 		if node == nil{
// 			return
// 		}
// 		res = append(res, node.Val)
// 		for _,v := range node.Children{
// 			pre(v)
// 		}
// 	}
// 	pre(root)
// 	return
// }

type Node struct {
	Val      int
	Children []*Node
}
