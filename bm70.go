package main

func MinMoney(arr []int, aim int) int {
	dp := make([]int, aim+1)
	for i := 1; i <= aim; i++ {
		dp[i] = aim + 1
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 1; i <= aim; i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] <= i {
				dp[i] = min(dp[i], dp[i-arr[j]]+1)
			}
		}
	}
	if dp[aim] > aim {
		return -1
	}
	return dp[aim]
}
