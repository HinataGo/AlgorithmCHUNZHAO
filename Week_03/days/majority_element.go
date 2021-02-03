package days

// Boyer-Moore 投票算法
// 算法可以分为两个阶段：
//    对抗阶段：分属两个候选人的票数进行两两对抗抵消
//    计数阶段：计算对抗结果中最后留下的候选人票数是否有效
func majorityElement(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major
}
