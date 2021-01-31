package tue

import "sort"

// 优先队列
var arr []int

type hp struct {
	// IntSlice将Interface的方法附加到[] int，并按升序排序
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return arr[h.IntSlice[i]] > arr[h.IntSlice[j]]
}
func (h *hp) Push(v interface{}) {
	// 由于IntSlice是Interface方法,这里元素是int 类型,所以使用断言转化一下
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
