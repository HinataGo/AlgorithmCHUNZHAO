package tue

// 滑动窗口
// 暴力解法 双端队列/单调队列
// 1. 这个与优先队列类似,不同点在于,不用自己手动实现优先队列,
// 2. 并且,处理思想上加了一个限制条件,即窗口头部数字 的索引过期, 并且得是 当前元素索引 比窗口size 大
// 3. 它会自己弹出队尾(当新的数字大于队尾存储的索引对应nums数字时)
// 4. 这里容易产生一个疑惑,会误认为一个突然特别大的数来了,队尾始终不会小于后面的数
// 比如 [1,3,-1,-3,-4,-6,-6,7]  3 ,这里3 比后面一大截都大,为什么还能正确输出呢?
// 因为题目会在关键时刻,删除 过期的队首,,而且是先删除过期队首,在删除队尾,(顺序不要错)
// 这样3 到后面超出索引会主动被删除,因此后面的数组的top元素都有了保证
func maxSlidingWindow2(nums []int, k int) []int {
	if nums == nil {
		return []int{}
	}
	// 这里windows直接理解为queue,存数组元素的下标
	window := make([]int, 0)
	res := make([]int, 0)
	for i, x := range nums {
		// 当前遍历到的下标大于窗口的大小, 并且, 床第一个存的数组的下标小于 i-k 的差值,
		// 例如 i = 4, k = 3, 随后多了一个 ,第一个存的下标是 0 自然删除 第一个存的下标数据
		if i >= k && window[0] <= i-k {
			// pop
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] <= x {
			// 如果队尾对应的值不比待插入的值大，则一直弹出 队尾
			window = window[:len(window)-1]
		}
		window = append(window, i)
		// push 压栈,存入新得遍历到的数据
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
