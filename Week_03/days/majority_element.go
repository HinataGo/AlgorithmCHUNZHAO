package days

// 169. 多数元素 (多数元素是指在数组中出现次数 大于  n/2  的元素 , n是元素总数)
// Boyer-Moore 投票算法
// 算法可以分为两个阶段：
//    对抗阶段：分属两个候选人的票数进行两两对抗抵消
//    计数阶段：计算对抗结果中最后留下的候选人票数是否有效

//
// 如果候选人不是maj 则 maj,会和其他非候选人一起反对 会反对候选人,所以候选人一定会下台(maj==0时发生换届选举)
// 如果候选人是maj , 则maj 会支持自己，其他候选人会反对，同样因为maj 票数超过一半，所以maj 一定会成功当选

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
