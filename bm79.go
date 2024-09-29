package main

func Bm79Rob(nums []int) int {
	n := len(nums)
	dp1 := make([]int, n+1)
	dp2 := make([]int, n+1)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp1[1] = nums[0]
	dp2[1] = 0
	for i := 2; i <= n; i++ {
		if i < n {
			dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i-1])
		}
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i-1])
	}
	return max(dp1[n-1], dp2[n])
}
