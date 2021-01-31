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
