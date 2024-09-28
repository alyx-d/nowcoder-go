package main

import (
	"sort"
	"strings"
)

func Permutation(s string) []string {
	ans := make([]string, 0)
	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	var dfs func(int)
	dfs = func(idx int) {
		if idx+1 == len(str) {
			ans = append(ans, strings.Clone(string(str)))
		}
		dict := make(map[byte]int, 0)
		for i := idx; i < len(str); i++ {
			if _, ok := dict[str[i]]; ok {
				continue
			}
			dict[str[i]] = 0
			str[i], str[idx] = str[idx], str[i]
			dfs(idx + 1)
			str[i], str[idx] = str[idx], str[i]
		}
	}
	dfs(0)
	return ans
}
