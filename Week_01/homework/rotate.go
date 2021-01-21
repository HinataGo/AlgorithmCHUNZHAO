package homework

// 189. 旋转数组
// 暴力的额外数组解法
// func rotate(nums []int, k int)  {
// 	// 最暴力的解法
// 	newNums := make([]int, len(nums))
// 	for i, v := range nums {
// 		newNums[(i+k)%len(nums)] = v
// 	}
// 	copy(nums, newNums)
// }

// 翻转数组 ，解法
// nums = "----->-->"; k =3
// result = "-->----->";
//
// reverse "----->-->" we can get "<--<-----"
// reverse "<--" we can get "--><-----"
// reverse "<-----" we can get "-->----->"
// this visualization help me figure it out :)
func rotate(nums []int, k int) {
	// 取余 是为了防止数组越界， k > len(nums)
	// 并且取余不会影响院线的关系
	k %= len(nums)
	// 取余结果为0 时其实不用翻转
	if k != 0 {
		reverse(nums)
		reverse(nums[:k])
		reverse(nums[k:])
	}
}

// 翻转函数
// 数组反转 很easy ，头尾交换即可，随后向中间逼近
func reverse(a []int) {
	// i小于数组长度一半就好， 奇数个数组两边换 ，偶数对半换，怎么也超不了
	// 8 --> 4 ，4
	// 5 --> 2 1 2  | n < 2 (数组开头时 0 1 2 3 4 )， 所以要记住规律
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
