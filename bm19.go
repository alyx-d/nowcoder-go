package main

func FindPeekElement(nums []int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
