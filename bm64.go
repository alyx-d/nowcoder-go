package main

func MinCostClimbingStairs(cost []int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
