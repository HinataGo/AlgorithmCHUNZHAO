package Week_05

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	dfs(n, n, "", n)
	return res
}

func dfs(lRemain int, rRemain int, path string, n int) {
	if 2*n == len(path) {
		res = append(res, path)
		return
	}
	if lRemain > 0 {
		dfs(lRemain-1, rRemain, path+"(", n)
	}
	if lRemain < rRemain {
		dfs(lRemain, rRemain-1, path+")", n)
	}
}
