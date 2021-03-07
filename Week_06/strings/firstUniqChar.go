package strings

func firstUniqChar(s string) int {
	res := [26]int{}
	for _, v := range s {
		res[v-'a']++
	}
	for i, v := range s {
		if res[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

// map 计数 ,,两次遍历查找
// 数组计数 索引使用  s[i] - a 存储,,随后遍历查找 v = ...... 有些繁琐,,抛弃
