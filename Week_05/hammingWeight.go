package Week_05

// 191. 位1的个数
/*
      n                   n-1             n & (n-1)
    01011               01010             01011 & 01010 -> 01010  count -> 1
	01010               01001             01010 & 01001 -> 01000  count -> 2
	01000               00111             01001 & 01000 -> 00000  count -> 3
	此时 n == 0 break
*/
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		num &= num - 1
		count++
	}
	return count
}
