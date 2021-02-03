package thu

// bfs 模板 迭代 循环版本  bfs 用于搜索效果一般好些
//  思想 假设一个三层的二叉树 从root 开始
// queue(双端队列 head tail ) root 从tail 进 head 出
// 1. root 入队 ,从head 出队 [root 遍历完] level ++
// 2. 第二层, 将左右孩子分别入队 即为 l,r ,
// 然后再拿出 l ,这是对 l 进行遍历, 让的 l.l l.r 入队, 此时l完成遍历
// 然后再从queue拿出 r , 遍历r 的 r.l r.r 入队,此时 r 完成遍历
// 后续依次进行 啊,bfs so easy

// 本题写法稍有不同, 这题需要很规范的输出每层每层的数据
func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}
