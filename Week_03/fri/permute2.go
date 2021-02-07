package fri

func permute2(nums []int) [][]int {
	res = make([][]int, 0)
	backtrack2(nums, 0, len(nums))
	return res
}
func backtrack2(nums []int, start, n int) {
	if start == n {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		res = append(res, tmp)
		return
	}
	for i := start; i < n; i++ {
		nums[i], nums[start] = nums[start], nums[i]
		backtrack2(nums, start+1, n)
		nums[i], nums[start] = nums[start], nums[i]
	}
}
