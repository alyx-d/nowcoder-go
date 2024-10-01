package main

import "sort"

func Bm89Merge(intervals []*Interval) []*Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	ans := make([]*Interval, 0)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _, val := range intervals {
		if len(ans) == 0 || ans[len(ans)-1].End < val.Start {
			ans = append(ans, val)
		} else {
			ans[len(ans)-1].End = max(ans[len(ans)-1].End, val.End)
		}
	}
	return ans
}
