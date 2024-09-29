package main

import "strconv"

func Bm69Solve(nums string) int {
	if len(nums) == 0 || nums == "0" {
		return 0
	}
	a, b := 1, 1
	sum := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == '0' {
			b = 0
		}
		val, _ := strconv.Atoi(nums[i-1 : i+1])
		if val >= 10 && val <= 26 {
			sum = a + b
		} else {
			sum = b
		}
		a, b = b, sum
	}
	return sum
}
