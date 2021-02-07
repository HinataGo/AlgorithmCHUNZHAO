package sat

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right { // 实际上是不会跳出循环，当 left==right 时直接返回
		if nums[left] <= nums[right] { // 如果 [left,right] 递增，直接返回
			return nums[left]
		}
		// 快速求 mid 方法
		mid := left + (right-left)>>1
		if nums[left] <= nums[mid] { // [left,mid] 连续递增，则在 [mid+1,right] 查找
			left = mid + 1
		} else {
			right = mid // [left,mid] 不连续，在 [left,mid] 查找
		}
	}
	return -1
}

// 暴力,从头到尾遍历比较
// 二分,,因为是半有序的数据, 所以开始搜索 先比较mid l ,r
