package thu

// bfs
// 二叉树的最小深度
// bfs 的遍历思想就是弄一个queue 没一层塞满之后
// 清空queue ,随后level ++
// 关键怎么判断 什么时候level++
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	var count []int
	queue = append(queue, root)
	count = append(count, 1)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}
