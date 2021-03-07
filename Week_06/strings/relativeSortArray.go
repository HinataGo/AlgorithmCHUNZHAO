package strings

// 计数排序
// 空间优化:  找出最大那个数,给他加1 就是我们需要创建的计数 数据的size
// 算法设计 :  第一次遍历arr1 然后计数次数  下标识对应数字,然后值是对应的次数
// 第二次 创建res ,然后, 遍历arr2 append arr2[i]  然后 用这个arr2[i] 搜索arr1 ,按次数 append (append一次 需要--)
// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾
func relativeSortArray(arr1 []int, arr2 []int) []int {
	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	tmp := make([]int, upper+1)
	for _, v := range arr1 {
		tmp[v]++
	}

	res := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for ; tmp[v] > 0; tmp[v]-- {
			res = append(res, v)
		}
	}
	for v, freq := range tmp {
		for ; freq > 0; freq-- {
			res = append(res, v)
		}
	}
	return res
}
