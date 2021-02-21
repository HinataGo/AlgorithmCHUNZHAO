package wen

// 借鉴爬楼梯的优化思路
// 这里将二维数组转换成一维数组
// 使用一位转换叠加结果,考虑01 状态
// 迭代到最后一步完成结果
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := make([]int, n+1)
	res[0], res[1] = 0, nums[0]
	for i := 2; i <= n; i++ {
		res[i] = max(res[i-1], res[i-2]+nums[i-1])
	}
	return res[n]
}
