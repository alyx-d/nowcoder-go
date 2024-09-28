package main

func Bm61Solve(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	ans := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var dfs func(int, int, int, int)
	dfs = func(x, y, val, count int) {
		if x < 0 || y < 0 || x >= m || y >= n || matrix[x][y] <= val {
			ans = max(ans, count)
			return
		}
		dfs(x-1, y, matrix[x][y], count+1)
		dfs(x+1, y, matrix[x][y], count+1)
		dfs(x, y-1, matrix[x][y], count+1)
		dfs(x, y+1, matrix[x][y], count+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, -1, 0)
		}
	}
	return ans
}
