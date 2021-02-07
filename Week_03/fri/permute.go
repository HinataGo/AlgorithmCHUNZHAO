package fri

// 学习回溯算法 入门
// 排列问题
// 46 全排列

// 回溯法 , 搜索树- 排列组合
// 先从集合内抽出元素 ,构造root
// 在使用剩余元素进行排列
/*
				123抽数
		1           2           3
	  2    3     1     3     1    2
	  3    2     3     1     2    1
	fun(arr[0:n-1]) --> 取出元素 e + fun(arr[0:n-1 - e] )
	随后不断递归,最后出来的就是排列ok的数据 (这题元素间是互斥的,不能同时出现)
*/

// 回溯算法 ,用完之后得还回去,,也就是恢复状态
// nums: 原始列表
// pathNums: 路径上的数字 存状态结果,后续append到 res 中返回
// used: 是否访问过  检查排列的位置可以放的元素,,检查谁被用了,,谁没被用
// 最终结果 一层一个位置,,所以就思考n层n位置放置数据,因此返回数组是二维数组
var res [][]int

func permute(nums []int) [][]int {
	// 初始化
	var pathNums []int
	var used = make([]bool, len(nums))
	// 清空全局数组  清一下全局变量
	res = [][]int{}
	// 回溯基于深搜,,这里用辅助函数解决,一步思考一个方案,,不要一口气瞎写
	backtrack(nums, pathNums, used)
	return res
}

// 完成辅助的深搜函数
func backtrack(nums, pathNums []int, used []bool) {
	// 结束条件：走完了，也就是路径上的数字总数等于原始列表总数
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, pathNums)
		// 把本次结果追加到最终结果上
		res = append(res, tmp)
		return
	}
	// 数据还有一部分为存入,寻找一个未使用的数存入
	// 开始遍历原始数组的每个数字
	for i := 0; i < len(nums); i++ {
		// 检查是否访问过
		if !used[i] {
			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true
			// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			pathNums = append(pathNums, nums[i])
			backtrack(nums, pathNums, used)
			// 撤销刚才的选择，也就是恢复操作 pop
			pathNums = pathNums[:len(pathNums)-1]
			// 标记成未使用
			used[i] = false
		}
	}
}
