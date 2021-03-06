package thu

import "math"

// dfs

// 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minNums := math.MaxInt64
	if root.Left != nil {
		minNums = min(minDepth(root.Left), minNums)
	}
	if root.Right != nil {
		minNums = min(minDepth(root.Right), minNums)
	}
	return minNums + 1
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
