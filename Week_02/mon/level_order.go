package mon

var res [][]int

func levelOrder(root *Node) [][]int {
	res = [][]int{}
	if root == nil {
		return res
	}
	queue := []*Node{root}
	var level int
	for len(queue) > 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			if queue[i] != nil {
				res[level] = append(res[level], queue[i].Val)
				for _, n := range queue[i].Children {
					queue = append(queue, n)
				}
			}
		}
		queue = queue[count:]
		level++
	}
	return res
}
