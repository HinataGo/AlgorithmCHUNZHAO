package strings

// 贪心算法
// 判断回文字符串 ： 定义左右指针，初始时分别指向字符串的第一个字符和最后一个字符，
// 每次判断左右指针指向的字符是否相同，
// 如果不相同，则不是回文串；
// 如果相同，则将左右指针都往中间移动一位，
// 直到左右指针相遇，则字符串是回文串。
//
// 允许最多删除一个字符的情况下，同样可以使用双指针，通过贪心算法实现。
// 初始化两个指针 low 和 high 分别指向字符串的第一个字符和最后一个字符。
// 每次判断两个指针指向的字符是否相同，
// 如果相同，则更新指针，令 low = low + 1 和 high = high - 1  然后判断更新后的指针范围内的子串是否是回文字符串
//
// 如果两个指针指向的字符不同，则两个字符中必须有一个被删除，此时我们就分成两种情况
// 删除左指针对应的字符 留下子串 s[low + 1], s[low + 1], ..., s[high]
// 删除右指针对应的字符，留下子串 s[low], s[low + 1], ..., s[high - 1]
//
// 当这两个子串中至少有一个是回文串时，就说明原始字符串删除一个字符之后就以成为回文串
func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			// 删除右指针
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			// 删除左指针
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			// 有一个真就是真
			return flag1 || flag2
		}
	}
	// 全部通过即为 true
	return true
}
