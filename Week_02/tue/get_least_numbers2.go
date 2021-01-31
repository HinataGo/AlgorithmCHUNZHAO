package tue

import "container/heap"

func getLeastNumbers(arr []int, k int) []int {
	var res []int
	n := len(arr)
	if n == 0 || k <= 0 || n < k {
		return res
	}
	pq := &maxHeap{}
	for i := 0; i < k; i++ {
		heap.Push(pq, arr[i])
	}
	heap.Init(pq)
	for i := k; i < n; i++ {
		if pq.Len() > 0 && arr[i] < pq.Peek().(int) {
			heap.Pop(pq)
			heap.Push(pq, arr[i])
		}
	}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(pq).(int))
	}
	return res
}

// 基于go interface实现最大堆
type maxHeap []int

func (h *maxHeap) Len() int {
	return len(*h)
}

func (h *maxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *maxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 查看堆顶元素
func (h *maxHeap) Peek() interface{} {
	return (*h)[0]
}
