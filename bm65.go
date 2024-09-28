package main

func LCS(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return "-1"
	}
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	ans := ""
	for i, j := len(s1), len(s2); dp[i][j] >= 1; {
		if s1[i-1] == s2[j-1] {
			ans += string(s1[i-1])
			i--
			j--
		} else if dp[i-1][j] >= dp[i][j] {
			i--
		} else {
			j--
		}
	}
	reverse := func(s string) string {
		chars := []byte(s)
		for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
			chars[i], chars[j] = chars[j], chars[i]
		}
		return string(chars)
	}
	if len(ans) == 0 {
		return "-1"
	} else {
		return reverse(ans)
	}
}
