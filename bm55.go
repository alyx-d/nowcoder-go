package main

import "sort"

func Permute(num []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(num)
	var dfs func(int)
	dfs = func(idx int) {
		if idx+1 == len(num) {
			dst := make([]int, len(num))
			copy(dst, num)
			ans = append(ans, dst)
		}
		for i := 1; i < len(num); i++ {
			num[i], num[idx] = num[idx], num[i]
			dfs(idx + 1)
			num[i], num[idx] = num[idx], num[i]
		}
	}
	dfs(0)
	return ans
}
