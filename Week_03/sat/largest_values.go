package sat

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历不太会,用深度优先遍历解
// 深度遍历,带上level,第一次 直接append
// 第二次则update为最大值
// 最后返回结果
var res []int

func largestValues(root *TreeNode) []int {
	res = make([]int, 0)
	dfs(root, 1)
	return res
}
func dfs(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if level == len(res)+1 {
		res = append(res, node.Val)
	} else {
		// tmp := res[level-1]
		// res = res[:level - 1]
		// res = append(res,max(tmp,node.Val))
		// update 为最大值
		res[level-1] = max(res[level-1], node.Val)
	}
	dfs(node.Left, level+1)
	dfs(node.Right, level+1)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
