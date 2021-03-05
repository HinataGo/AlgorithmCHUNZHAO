package mon

func maxArea(height []int) int {
	area := 0
	l, r := 0, len(height)-1
	for l != r {
		hl, hr := height[l], height[r]
		s := (r - l) * min(hl, hr)
		if s > area {
			area = s
		}

		if hl > hr {
			r--
		} else {
			l++
		}
	}
	return area
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
