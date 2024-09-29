package main

func LIS(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	dp := make([]int, len(arr))
	maxLen := 1
	dp[0] = 1
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}
