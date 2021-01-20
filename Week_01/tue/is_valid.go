package tue

//  lc 20. 有效的括号
func isValid(s string) bool {
	// 常见解法，用一个stack 去匹配数据，pop 和 push
	// string 的每个字符单位是 rune
	n := len(s)
	if n%2 != 0 {
		return false
	}
	stack := make([]rune, 0)
	table := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		if table[v] != 0 {
			if len(stack) == 0 || stack[len(stack)-1] != table[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}

// stack 其实用 interface{}最好，但是这里题目要求 string 参数类型， 所以使用[]string  数组
// 这里还要注意一个点，因为时再题目目中写，stack 只需要自行模拟一个简单的即可
