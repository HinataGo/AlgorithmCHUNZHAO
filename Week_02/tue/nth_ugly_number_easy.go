package tue

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for num != 1 {
		switch {
		case num&0x01 == 0:
			num = num >> 1
		case num%3 == 0:
			num = num / 3
		case num%5 == 0:
			num = num / 5
		default:
			return false
		}
	}
	return true
}
