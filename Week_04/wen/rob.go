package wen

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := make2Slice(n, 2)
	res[0][0] = 0
	res[0][1] = nums[0]
	for i := 1; i < n; i++ {
		res[i][1] = res[i-1][0] + nums[i]
		res[i][0] = max(res[i-1][0], res[i-1][1])

	}
	return max(res[n-1][0], res[n-1][1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func make2Slice(size int, n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}

// 0 不偷 1  偷

// res[0][0] = 0
// res[0][1] = nums[0]

// res[i][0] = max(res[i-1][0],res[i-1][1])
// res[i][1] = res[i-1][0] + nums[i]
