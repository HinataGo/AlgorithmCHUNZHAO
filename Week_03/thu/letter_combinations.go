package thu

// 17. 电话号码的字母组合
var table = []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
var res []string

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	if digits == "" {
		return res
	}
	find(digits, 0, "")
	return res
}
func find(digits string, index int, tmp string) {
	// 终止条件
	if index == len(digits) {
		res = append(res, tmp)
		return
	}
	// drill down
	v := digits[index]
	letter := table[v-'0']
	for i := range letter {
		find(digits, index+1, tmp+string(letter[i]))
	}
}
