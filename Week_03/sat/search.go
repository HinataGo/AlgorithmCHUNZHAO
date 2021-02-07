package sat

// 二分模板题,但是这里因为划分为两段有序,所以比较复杂
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if (nums[l] < nums[mid] && target > nums[r] && target < nums[mid]) ||
			(nums[mid] < nums[r] && (target < nums[mid] || target > nums[r])) {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return -1
}
