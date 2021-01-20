package wen

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}
