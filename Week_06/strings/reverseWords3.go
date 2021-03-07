package strings

// 遇到 空， 反转 l-- i 的数据即可
// 单独定义反转函数 对称交换
func reverseWords3(s string) string {
	b := []byte(s)
	l := -1
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			reverse(b[l+1 : i])
			l = i
		}
	}
	reverse(b[l+1:])
	return string(b)
}

func reverse(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
}
