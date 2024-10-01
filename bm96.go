package main

import "sort"

func MinmumNumberOfHost(n int, startEnd [][]int) int {
	ans := 0
	start := make([]int, len(startEnd))
	end := make([]int, len(startEnd))
	for idx, val := range startEnd {
		start[idx] = val[0]
		end[idx] = val[1]
	}
	sort.Ints(start)
	sort.Ints(end)
	for i, j := 0, 0; i < len(startEnd); i++ {
		if start[i] >= end[j] {
			j++
		} else {
			ans++
		}
	}
	return ans
}
