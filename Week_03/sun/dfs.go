package sun

// dfs 模板
// 前中后
var ret []int

func dfs(root *TreeNode) []int {
	ret = make([]int, 0)
	pre(root)
	return ret
}

func pre(node *TreeNode) {
	if node == nil {
		return
	}
	// 下面三个分别代表, 根 左 右,
	// 其他的遍历只需要调换这三者之间的顺序即可
	ret = append(ret, node.Val)
	pre(node.Left)
	pre(node.Right)
}

// 迭代版本, 自己维护一个栈
// 遍历 len(stack) > 0  || root != nil
func preorderStack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	// root != nil 只为了第一次root判断，必须放最后
	for 0 < len(stack) || root != nil {
		for root != nil {
			res = append(res, root.Val)       // 前序遍历
			stack = append(stack, root.Right) // 右节点 入栈
			root = root.Left                  // 移至最左
		}
		index := len(stack) - 1 // 栈顶
		root = stack[index]     // 出栈
		stack = stack[:index]
	}
	return res
}

// 中序遍历迭代
func inorderStack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil {
		for root != nil {
			// 入栈
			stack = append(stack, root)
			// 移至最左
			root = root.Left
		}
		index := len(stack) - 1             // 栈顶
		res = append(res, stack[index].Val) // 中序遍历
		root = stack[index].Right           // 右节点会进入下次循环，如果 =nil，继续出栈
		stack = stack[:index]               // 出栈
	}
	return res
}

// 后续遍历迭代
// 在前序遍历基础上加一层翻转
// 翻转方法 l,r := 0, len -1
// for l < r  --> swap(l,r)
func postorderStack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	for 0 < len(stack) {
		if root != nil {
			// 根右左
			res = append(res, root.Val)
			stack = append(stack, root.Left)  // 左节点，因为先进 所以后出
			stack = append(stack, root.Right) // 右节点，因为后进 所以先出
		}
		index := len(stack) - 1 // 栈顶
		root = stack[index]     // 出栈
		stack = stack[:index]
	}
	// 反转 变成后序遍历 左右根
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
