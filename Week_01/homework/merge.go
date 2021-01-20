package homework

func merge(A []int, m int, B []int, n int) {
	i, j := m-1, n-1
	k := len(A) - 1

	for i >= 0 && j >= 0 {
		if A[i] > B[j] {
			A[k] = A[i]
			i--
		} else {
			A[k] = B[j]
			j--
		}
		k--
	}
	// 等于0,  特点情况
	// [0]
	// 0
	// [1]
	// 1
	if j >= 0 {
		copy(A[:j+1], B[:j+1])
	}
}
