package Week_06

import "sort"

// prev 初始为第一个区间，cur 表示当前的区间，res 表示结果数组
//
//    开启遍历，尝试合并 prev 和 cur，合并后更新到 prev 上
//    因为合并后的新区间还可能和后面的区间重合，继续尝试合并新的 cur，更新给 prev
//    直到不能合并 —— prev[1] < cur[0]，此时将 prev 区间推入 res 数组
//
// 合并的策略
//
//    原则上要更新prev[0]和prev[1]，即左右端:
//        prev[0] = min(prev[0], cur[0])
//        prev[1] = max(prev[1], cur[1])
//    但如果先按区间的左端排升序，就能保证 prev[0] < cur[0]
//    所以合并只需这条：prev[1] = max(prev[1], cur[1])
//
//  我们是先合并，遇到不重合再推入 prev。
// 当考察完最后一个区间，后面没区间了，遇不到不重合区间，最后的 prev 没推入 res。
// 要单独补上
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { // 没有一点重合
			res = append(res, prev)
			prev = cur
		} else { // 有重合
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
