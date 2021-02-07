package sat

import "math"

func largestValues2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var length = len(queue)
		var max = math.MinInt64 // -1 << 63
		for i := 0; i < length; i++ {
			// 取max值
			// if max < queue[i].Val {
			// 	max = queue[i].Val
			// }
			max = getMax(max, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, max)
		queue = queue[length:]
	}
	return res
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
