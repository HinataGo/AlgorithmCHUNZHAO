package Week_06

func lengthOfLIS(nums []int) int {
	ans := 0
	if n := len(nums); n > 0 {
		d := make([]int, n)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if nums[j] > nums[i] {
					if d[j] < d[i]+1 {
						d[j] = d[i] + 1
						if d[j] > ans {
							ans = d[j]
						}
					}
				}
			}
		}
		ans++
	}
	return ans
}
