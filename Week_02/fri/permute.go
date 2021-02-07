package fri

func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(nums []int, start, n int)
	backtrack = func(nums []int, start, n int) {
		if start == n {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}
		for i := start; i < n; i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backtrack(nums, start+1, n)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	backtrack(nums, 0, len(nums))
	return res
}
