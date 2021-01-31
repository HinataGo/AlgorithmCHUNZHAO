package tue

import "container/heap"

// 使用golang 自定义heap 实现-->最大堆
// 根据最大堆的性质,遍历前k个元素 ,pop即可,注意不是peek ,斗则数据就错了
//
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
