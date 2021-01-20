package mon

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	if n <= 1 {
		return 0
	}
	res := (r - l) * min(height[l], height[r])
	for l < r {
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		tmp := (r - l) * min(height[l], height[r])
		res = max(tmp, res)
	}
	return res
}

func max(tmp int, res int) int {
	if tmp > res {
		return tmp
	}
	return res
}

func min(i int, i2 int) int {
	if i > i2 {
		return i2
	}
	return i
}
