package wen

import "sort"

// 排序思路
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 很简单，字符串转 字符数组 ，内部排序， 转字符串对比
func isAnagram3(s string, t string) bool {
	s1 := []byte(s)
	t1 := []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(t1, func(i, j int) bool { return t1[i] < t1[j] })
	return string(s1) == string(t1)
}
