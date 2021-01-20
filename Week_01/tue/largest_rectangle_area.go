package tue

// 暴力解法
// 双层遍历
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	res := 0
	for i, v := range heights {
		if v < heights[i-1] && i > 0 {

		}
		if v > heights[i-1] && i < n-1 {

		}
	}
	return res
}
