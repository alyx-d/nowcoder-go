package main

import "sort"

func PermuteUnique(num []int) [][]int {
	ans := make([][]int, 0)
	var dfs func(int)
	dfs = func(idx int) {
		if idx+1 == len(num) {
			dst := make([]int, len(num))
			copy(dst, num)
			ans = append(ans, dst)
		}
		dict := make(map[int]int, 0)
		for i := idx; i < len(num); i++ {
			if _, ok := dict[num[i]]; ok {
				continue
			}
			dict[num[i]] = 0
			num[i], num[idx] = num[idx], num[i]
			dfs(idx + 1)
			num[i], num[idx] = num[idx], num[i]
		}
	}
	sort.Ints(num)
	dfs(0)
	return ans
}
