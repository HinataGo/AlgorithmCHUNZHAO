package tue

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
	return arr[:k]
}

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
