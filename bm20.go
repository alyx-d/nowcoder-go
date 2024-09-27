package main

var ans = 0
var mod = 1000000007

func InversePairs(nums []int) int {
	mergeSort(nums, 0, len(nums)-1)
	return ans
}

func mergeSort(nums []int, left int, right int) []int {
	if left >= right {
		return nums
	}
	mid := (left + right) >> 1
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	return mergeRange(nums, left, mid, right)
}

func mergeRange(nums []int, left int, mid int, right int) []int {
	length := right - left + 1
	temp := make([]int, length)
	i := left
	j := mid + 1
	k := 0
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
			ans += mid - i + 1
			ans %= mod
		}
		k++
	}
	for i <= mid {
		temp[k] = nums[i]
		k++
		i++
	}
	for j <= right {
		temp[k] = nums[j]
		k++
		j++
	}
	k = 0
	i = left
	for ; k < length; k++ {
		nums[i] = temp[k]
		i++
	}
	return nums
}
