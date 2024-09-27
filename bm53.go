package main

func MinNumberDisappeared(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		if num < 0 {
			nums[i] = n + 1
		}
	}
	for _, num := range nums {
		k := abs1(num)
		if k >= 1 && k <= n {
			nums[k-1] = -abs1(nums[k-1])
		}
	}
	i := 0
	for i < n {
		if nums[i] > 0 {
			break
		}
		i++
	}
	return i + 1
}

func abs1(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
