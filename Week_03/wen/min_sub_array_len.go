package wen

// 个人练习
// 209. 长度最小的子数组
// 这里要明确滑动窗口的意义,随后有指针开始定义要和左指针不一样,r = -1
// 然后看需要移动++右指针(),然后给sum加上自己的值 ,确保新得值加入进来
// 先减去l自己的值 随后在 l++ 防止数组越界
// 第三步 才是判断两者额数据 sum >= s,
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	//
	l, r, sum, res := 0, -1, 0, n+1
	for l < n {
		if r+1 < n && sum < s {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s {
			res = min(res, r-l+1)
		}
	}
	if res == n+1 {
		return 0
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
