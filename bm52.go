package main

func FindNumsAppearOnce(nums []int) []int {
	s := 0
	for _, num := range nums {
		s ^= num
	}
	k := 1
	for k&s == 0 {
		k <<= 1
	}
	a, b := 0, 0
	for _, num := range nums {
		if k&num == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	if a < b {
		return []int{a, b}
	} else {
		return []int{b, a}
	}
}
