package tue

// 堆解法  ,推荐动态规划,太慢了这个,还麻烦
// 时间复杂度：O(1) 的时间检索答案,但是有 超过 10^612×106 次预计算操作
// 空间复杂度：常数空间存储 1690 个丑数，和堆中不超过  21690×2 的元素和
// 哈希表中不超过  31690×3 的元素

// 小顶堆的方法是先存再排，dp的方法则是先排再存
// 首先把丑数初始化为一个不合法的值-1
// 然后定义最小堆
// 并且把第一个丑数1，加入到最小堆中
// 当n大于0时执行以下操作
// 检查堆顶元素是否等于上一个丑数
// 如果是就不断丢弃
// 这一步的作用是去重
// 直到不重复就取出堆顶元素,把它作为新的丑数
// 接着把当前丑数乘以2,3,5后的候选值，加入最小堆中
// 由于候选集合增长非常快，有可能在还没有求出第n个丑数时，
// 就已经超出了整数最大值，于是这里使用整数的最大值进行约束
// 计算出一个新的丑数后n要减一

import (
	"container/heap"
	"math"
)

func nthUglyNumber2(n int) int {
	ugly := -1
	var mh minHeap
	heap.Init(&mh)
	heap.Push(&mh, 1)
	for n > 0 {
		for mh[0] == ugly {
			heap.Pop(&mh)
		}
		ugly = heap.Pop(&mh).(int)

		if 2*ugly <= math.MaxInt32 {
			heap.Push(&mh, 2*ugly)
		}
		if 3*ugly <= math.MaxInt32 {
			heap.Push(&mh, 3*ugly)
		}
		if 5*ugly <= math.MaxInt32 {
			heap.Push(&mh, 5*ugly)
		}
		n--
	}
	return ugly
}
