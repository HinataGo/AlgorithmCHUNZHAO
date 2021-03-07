package strings

// 简化后版本
// 1.中心扩散，先找出和中心不相等的左、右索引，这样可以忽略奇、偶长度对称性的差异
// 2.下一次搜索直接将索引移动到上一次的初始右索引，节省重复计算
// Todo:其实全部循环完成后再生成子字符串就行了
//
func longestPalindrome(s string) string {
	var L, R = 0, 0
	var ansL, ansR = 0, 0
	for index := 0; index < len(s)-1; {
		L, R = index, index
		for ; L >= 0 && s[L] == s[index]; L-- {
		}
		for ; R < len(s) && s[R] == s[index]; R++ {
		}
		index = R
		for ; L >= 0 && R < len(s) && s[R] == s[L]; L, R = L-1, R+1 {
		}
		if R-L-2 > ansR-ansL {
			ansL, ansR = L+1, R-1
		}
	}
	return string(s[ansL : ansR+1])
}

// func longestPalindrome(s string) string {
// 	var index int = 0
// 	var size int = len(s)
// 	var maxLen = 0
// 	var res = ""
// 	for index < size{
// 		var pointerL = index
// 		var pointerR = index
// 		for pointerL >= 0 && s[pointerL] == s[index]{
// 			pointerL--
// 		}
// 		for pointerR < size && s[pointerR] == s[index]{
// 			pointerR++
// 		}
// 		var nextPoint = pointerR
// 		for pointerL >= 0 && pointerR < size && s[pointerR] == s[pointerL]{
// 			pointerL--
// 			pointerR++
// 		}
// 		var tmpMaxLen = pointerR - pointerL - 1
// 		if tmpMaxLen > maxLen{
// 			res = s[pointerL + 1:pointerR]
// 			maxLen = tmpMaxLen
// 		}
//
// 		index = nextPoint
// 	}
// 	return res
// }
