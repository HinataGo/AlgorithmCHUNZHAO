package mon

// 原地操作
// 遍历交换
// 双指针 , l 0  ,r !0  swap
// 一定要同时从0 开始 , 防止 连续两个不为0 的数交换,,乱次序
// 让l 先 r 一步,随后只有当 nums[r] == 0 ,两者才会错开, 否则两者同步
// 交换自然不会出错

func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}

}
