package Week_06

func lengthOfLIS(nums []int) int {
	d := make([]int, 0)
	l := 1
	d = append(d, nums[0])
	for i := 0; i < len(nums); i++ {
		if nums[i] > d[l-1] {
			d = append(d, nums[i])
			l = l + 1
		} else {
			get(d, nums[i])
		}
	}
	return len(d)
}

// d是一个单调递增序列
// 获取数组d中第一个小于a的值d[k]，并赋值d[k+1]=a
func get(d []int, a int) {
	for i := 0; i < len(d); i++ {
		if d[i] >= a {
			d[i] = a
			break
		}
	}
}
