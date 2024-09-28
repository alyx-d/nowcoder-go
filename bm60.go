package main

func GenerateParenthesis(n int) []string {
	ans := make([]string, 0)
	var dfs func(int, int, string)
	dfs = func(l, r int, temp string) {
		if l == n && r == n {
			ans = append(ans, temp)
			return
		}
		if l < n {
			dfs(l+1, r, temp+"(")
		}
		if r < n && l > r {
			dfs(l, r+1, temp+")")
		}
	}
	dfs(0, 0, "")
	return ans
}
