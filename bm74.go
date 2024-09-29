package main

import "strconv"

func RestoreIpAddresses(s string) []string {
	nums := ""
	ans := make([]string, 0)
	n := len(s)
	var dfs func(int, int)
	dfs = func(step, idx int) {
		cur := ""
		if step == 4 {
			if idx != n {
				return
			}
			ans = append(ans, nums)
			return
		}
		for i := idx; i < idx+3 && i < n; i++ {
			cur += string(s[i])
			num, _ := strconv.Atoi(cur)
			temp := nums
			if num <= 255 && (len(cur) == 1 || cur[0] != '0') {
				if step < 3 {
					nums += cur + "."
				} else {
					nums += cur
				}
				dfs(step+1, i+1)
				nums = temp
			}
		}
	}
	dfs(0, 0)
	return ans
}
