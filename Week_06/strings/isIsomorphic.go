package strings

// 首先判断映射
// 建立映射，

func isIsomorphic(s string, t string) bool {
	sFirst, tFirst := make([]int, 128), make([]int, 128)

	for i := 0; i < len(s); i++ {
		if sFirst[s[i]] != tFirst[t[i]] {
			return false
		}
		sFirst[s[i]] = i + 1
		tFirst[t[i]] = i + 1
	}
	return true
}
