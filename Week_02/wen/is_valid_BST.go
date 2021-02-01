package wen

import "math"

// 方法1. BST 中序遍历数据递增
// 方法2. 写一个递归
// 综合一下,遍历过程中就进行对比, 递归大法好!

var pre = math.MinInt64

func isValidBST(root *TreeNode) bool {
	// 再次初始化全局变量,防止数据被修改后出现问题
	pre = math.MinInt64
	return is(root)
}

func is(root *TreeNode) bool {
	// 基本条件检查
	if root == nil {
		return true
	}
	// 访问左子树
	if !is(root.Left) {
		return false
	}
	// 中序遍历的升序验证
	if root.Val <= pre {
		return false
	}
	//  更新前一个节点值
	pre = root.Val
	// 访问右子树
	return is(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
