package mon

// 28. 实现 strStr()
// rabin-karp
// hash func
// str1 : abcde  str2: bcd
// (a*31^4 + b*31^3 ...+e*31^0 ) % 10^6
// hash 的问题是 hash后的数据一样,,但是原串可能不同
// 这是在字符串比较一下

var base = 1000000

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	n := len(haystack)
	m := len(needle)

	power := 1
	for i := 0; i < m; i++ {
		power = (power * 31) % base
	}
	targetCode := 0
	for i := 0; i < m; i++ {
		targetCode = (targetCode*31 + int(needle[i])) % base
	}
	hashCode := 0
	for i := 0; i < n; i++ {
		// abc + d
		hashCode = (hashCode*31 + int(haystack[i])) % base
		if i < m-1 {
			continue
		}
		// abcd - a
		if i >= m {
			hashCode = (hashCode - (power * int(haystack[i-m]))) % base
			if hashCode < 0 {
				hashCode += base
			}
		}
		if hashCode == targetCode {
			if haystack[i-m+1:i+1] == needle {
				return i - m + 1
			}
		}
	}
	return -1
}
