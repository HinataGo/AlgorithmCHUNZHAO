package strings

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	table := make(map[rune]int, 0)
	for _, v := range s {
		table[v]++
	}
	for _, v := range t {
		if table[v] != 0 {
			table[v]--
		} else {
			return false
		}
	}
	return true
}
