package wen

// hash思路
// 1.确定26字母时的做法
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var a1, a2 [26]rune
	for i := range s {
		a1[s[i-'a']]++
	}
	for i := range t {
		a2[t[i-'a']]++
	}
	return a1 == a2
}
