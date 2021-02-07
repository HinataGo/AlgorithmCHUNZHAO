package fri

func subsets(nums []int) (res [][]int) {
	var dfs func(index int, list []int)
	dfs = func(index int, list []int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)

		for i := index; i < len(nums); i++ {
			list = append(list, nums[i])
			dfs(i+1, list)
			list = list[:len(list)-1]
		}
	}
	dfs(0, []int{})
	return
}
