package strings

// 反转函数
// 两头开工，通过是交换 l ，r
// l < r  ！is(l) l++
// l < r  ！is(r) r--
// swap(l,r)
// l++ r--
func reverseOnlyLetters(S string) string {
	buf := []byte(S)
	for l, r := 0, len(buf)-1; l < r; {
		for l < r && !isLetter(buf[l]) {
			l++
		}
		for l < r && !isLetter(buf[r]) {
			r--
		}
		buf[l], buf[r] = buf[r], buf[l]
		l++
		r--
	}
	return string(buf)
}

// 判断是否是字符
func isLetter(c byte) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
