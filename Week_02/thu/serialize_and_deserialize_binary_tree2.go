package thu

import (
	"strconv"
	"strings"
)

// DFS 解法
type Codec struct {
	s []string
}

func Constructor() Codec {
	return Codec{}
}

// 先序遍历
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," +
		this.serialize(root.Left) + "," +
		this.serialize(root.Right)
}

// 这里必须用指针,否则数据无法修改,会导致错误
func dfsTree(str *[]string) *TreeNode {
	val := (*str)[0]
	*str = (*str)[1:]
	if val == "#" {
		return nil
	}
	valInt, _ := strconv.Atoi(val)
	node := &TreeNode{Val: valInt}
	node.Left = dfsTree(str)
	node.Right = dfsTree(str)
	return node
}

// 反序列化,就是将这个之前的遍历一个一个在解析出来,
// 相当于再次根据数据信息,自己建树,前后顺序保持一致
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return dfsTree(&vals)
}
