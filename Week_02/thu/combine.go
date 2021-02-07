package thu

// lc 77. 组合
// 递归实现组合型枚举
var ret [][]int

func combine(n int, k int) [][]int {
	ret = make([][]int, 0)
	if n <= 0 || k <= 0 || k > n {
		return ret
	}
	tmp := make([]int, 0)
	find(n, k, 1, tmp)
	return ret
}

// 所有可能的 k 个数的组合
func find(n, k int, start int, tmp []int) {
	// 不合法参数
	if len(tmp) == k {
		// 这里不能直接append,否则数据错误
		path := make([]int, len(tmp))
		copy(path, tmp)
		ret = append(ret, path)
		return
	}
	// 这里优化 剪枝 最后 i== n的枝丫不需要
	// n --> n
	for i := start; i <= n-(k-len(tmp))+1; i++ {
		tmp = append(tmp, i)
		find(n, k, i+1, tmp)
		// 恢复状态
		tmp = tmp[:len(tmp)-1]
	}
}
