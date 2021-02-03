package thu

import (
	"strconv"
	"strings"
)

// 二叉树的序列化与反序列化 这题难度相当大
// 官网题解,很g hen ex
// 时间复杂度：在序列化和反序列化函数中，我们只访问每个节点一次，
// 因此时间复杂度为 O(n)，其中 nnn 是节点数，即树的大小。
// 空间复杂度：在序列化和反序列化函数中，我们递归会使用栈空间，
// 故渐进空间复杂度为 O(n)

// TODO
// 这里使用 bfs 解决
func (this *Codec) serialize1(root *TreeNode) string {
	stack := []*TreeNode{root}
	vals := make([]string, 0)
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			vals = append(vals, "#")
		} else {
			vals = append(vals, strconv.Itoa(node.Val))
			stack = append(stack, node.Left, node.Right)
		}
	}
	return strings.Join(vals, ",")
}

func (this *Codec) deserialize1(data string) *TreeNode {
	vals := strings.Split(data, ",")
	idx := 0
	if vals[idx] == "#" {
		return nil
	}
	val, _ := strconv.Atoi(vals[idx])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	var node, left, right *TreeNode
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			left = &TreeNode{Val: val}
			node.Left = left
			queue = append(queue, left)
		}
		idx++
		if vals[idx] != "#" {
			val, _ = strconv.Atoi(vals[idx])
			right = &TreeNode{Val: val}
			node.Right = right
			queue = append(queue, right)
		}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
