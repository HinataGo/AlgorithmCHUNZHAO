package UnionFInd

var p []int

func findCircleNum(M [][]int) int {
	n := len(M)
	p = make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	ans := n
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if M[i][j] == 1 && find(i) != find(j) {
				p[find(i)] = find(j)
				ans--
			}
		}
	}
	return ans
}

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}
