package mon

func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		j := 0
		for ; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}

		if j == len(needle) {
			return i
		}

	}

	return -1

}
