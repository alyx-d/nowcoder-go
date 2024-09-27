package main

func GetLeastNumbers_Solution(input []int, k int) []int {
	l, r := 0, len(input)-1
	for l <= r {
		pi := partition2(input, l, r)
		if pi+1 == k {
			return input[0:k]
		} else if pi+1 < k {
			l = pi + 1
		} else {
			r = pi - 1
		}
	}
	return make([]int, 0)
}

func partition2(nums []int, p, q int) int {
	pivot := nums[q]
	i := p - 1
	for j := p; j <= q-1; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[q] = nums[q], nums[i+1]
	return i + 1
}
