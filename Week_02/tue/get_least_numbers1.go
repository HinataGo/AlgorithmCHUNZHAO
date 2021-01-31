package tue

// 快排 (升序思想)
// 下面的是第二部分快排思路
// 定义好 左右指针 和 index 数轴值,
// 由于数轴是可以划分数据的 , 数轴左边皆是小,右边皆大
// 所以top k 看index == k 并且返回后面的数就好了
// for index != k 时不断更新index 值, 找到合适的数轴
// index > k-1(一定要减去一,下标从零开始的问题) 大了,说明左边有一部分数据会遗漏, 就把index 缩小,给r
// index < k-1 , 小了,说明自己多拿了一些数据,因此缩小范围,index 右移动
// 至此,本地 ac
func getLeastNumbers1(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	l, r := 0, len(arr)-1
	index := qs(arr, l, r)
	for index != k-1 {
		if index > k-1 {
			r = index - 1
		} else {
			l = index + 1
		}
		index = qs(arr, l, r)
	}
	// 只有这一行是返回数据top k
	return arr[:k]
}

// 再写一次快排思路加深印象 (最简单的快排)
// 定 枢轴 e,左右指针 l,r
// for l < r -->
// 下面目的是为了遍历数据,往 数轴 e 靠拢
// 判断右指针值 arr[r] > e 有指针左移  减小 --> 把 有指针值 赋值给 左指针,不用做交换
// 左指针 arr[l] < e 左指针右移 增大 --> 把左指针值赋值给右指针
// 最后返回数轴给左指针
// 返回 l 的值即可,也可以不返回
func qs(arr []int, l, r int) int {
	e := arr[l]
	for l < r {
		for l < r && arr[r] >= e {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] < e {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = e
	return l
}
