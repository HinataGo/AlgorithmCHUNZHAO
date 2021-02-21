package wen

// 53. 最大子序和

// func maxSubArray(nums []int) int {
// 	n := len(nums)
// 	max := nums[0]
// 	for i := 1; i < n; i++ {
// 		if nums[i-1]+nums[i] > nums[i] {
// 			nums[i] += nums[i-1]
// 		}
// 		if nums[i] > max {
// 			max = nums[i]
// 		}
// 	}
// 	return max
// }

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0

	}
	res := nums[0]
	for i := 1; i < n; i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		res = max(res, nums[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// f(n) = f(n-1) + n
// 做一个easy判断, 正向移动遍历, 到达i时 如果 i-1 + i > i 那么加上之前的值
// 否则 跳过,
// 如果当前值大于max , 那么则更新当前的值
