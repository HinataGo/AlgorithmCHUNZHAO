package Week_06

// 空间优化
func findAnagrams(s string, p string) (res []int) {
	needs := [26]int{}
	window := [26]int{}
	match := 0
	left, right := 0, 0
	needsLen := 0
	for _, c := range p {
		if needs[byte(c)-'a'] == 0 {
			needsLen++
		}
		needs[byte(c)-'a']++
	}

	for right < len(s) {
		if needs[s[right]-'a'] > 0 {
			window[s[right]-'a']++
			if window[s[right]-'a'] == needs[s[right]-'a'] {
				match++
			}
		}
		right++
		for match == needsLen {
			if right-left == len(p) {
				res = append(res, left)
			}
			if needs[s[left]-'a'] > 0 {
				window[s[left]-'a']--
				if window[s[left]-'a'] < needs[s[left]-'a'] {
					match--
				}
			}
			left++
		}
	}
	return
}

// func findAnagrams(s string, p string) (res []int) {
// 	needs := make(map[byte]int)
// 	window := make(map[byte]int)
// 	match := 0
// 	left, right := 0, 0
// 	for _, c := range p {
// 		needs[byte(c)]++
// 	}
// 	for right < len(s) {
// 		if _, ok := needs[s[right]]; ok {
// 			window[s[right]]++
// 			if window[s[right]] == needs[s[right]] {
// 				match++
// 			}
// 		}
// 		right++
// 		for match == len(needs) {
// 			if right-left == len(p) {
// 				res = append(res, left)
// 			}
// 			if _, ok := needs[s[left]]; ok {
// 				window[s[left]]--
// 				if window[s[left]] < needs[s[left]] {
// 					match--
// 				}
// 			}
// 			left++
// 		}
// 	}
// 	return
// }
