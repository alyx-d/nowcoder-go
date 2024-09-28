package main

func Bm66LCS(s1, s2 string) string {
	maxLen := 0
	pos := 0
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			if maxLen < dp[i][j] {
				maxLen = dp[i][j]
				pos = i - 1
			}
		}
	}
	return s1[pos-maxLen+1 : pos+1]
}
