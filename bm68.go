package main

func MinPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 {
				dp[i][j] += dp[i][j-1] + matrix[i-1][j-1]
				continue
			}
			if j == 1 {
				dp[i][j] += dp[i-1][j] + matrix[i-1][j-1]
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i-1][j-1]
		}
	}
	return dp[m][n]
}
