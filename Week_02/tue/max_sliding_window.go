package tue

import "container/heap"

// 优先队列
// golang continuer 包自带 heap
// 自己实现一个最大堆
// 时间复杂度 O(nlogn)
// 空间复杂度O(n)

// 思想: 有一个优先队列,队首始终放最大的数字
// 然后这个队列开始append数字, 新的来了后,
// 1. 如果当前最大的数字是size之外的 即 max_index(最大元素在队列中的下标) < now_index(当前遍历nums到的元素对应下标) - k(队列大小)
// 这是无论如何必须删除 max_index ,也就是删去队首
// 2. 如果最大的不是待删除的数字,那么新的数字一来(新数字)也不能删,
// 只能想办法删除那个超过size的数字(注意这里它可能既不是最大,也不是最小),这样思考我们会陷入一个寻找数字极其困难的情况
// 所以,我们干脆不去删除
// 3. 因此我们只需要,遇到队首的元素,不在范围时(队首的index < i - k)才删除(队首首先是最大的,其次我们不想那个特殊的元素了的元素的)
// 注意这里queue 的size 可能超过 k 但是不必担心,因为这里只需要防止最大树的混乱, 目标只是求对应窗口的最大值,最终整合成一个数组

// code 思路就是
// 1. 创建一个优先队列 使用golang 的sort.IntSlice 方便快捷,
// 后续记得初始化才能使用(初始化先给前k 个元素push到 queue中,并且使用sort.IntSlice排序)
// 2. 遍历nums , 遇到数字就开始push 进队列,
// 3. 然后,开始在队列中寻找,看队首数字是否已经过期(认为,窗口左边已经圈过,但是现在不在的的数字)
// 4. 将队首,append 给res 数组,最后只返回数组即可
func maxSlidingWindow(nums []int, k int) []int {
	arr = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)
	n := len(nums)
	res := make([]int, 1, n-k+1)
	res[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		res = append(res, nums[q.IntSlice[0]])
	}
	return res
}
