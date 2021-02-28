package Week_05

// dp 因为n是正整数 从 1 开始，随用用这种最高效的
func climbStairs(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}
	return b
}

// f0 1
// f1 1
// f2 2
// f3 f2 + f1
