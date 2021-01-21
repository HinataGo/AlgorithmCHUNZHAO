package wen

// 这题一定要问清楚字符范围，至关重要
// 针对Unicode字符集的做法，
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	res := make(map[rune]int, 0)
	for _, v := range s {
		if res[v] != 0 {
			res[v]++
		}
	}
	for _, v := range t {
		if res[v] != 0 {
			res[v]--
		} else {
			return false
		}
	}
	return true
}
