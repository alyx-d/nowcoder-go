package sort

import "math/rand"

func partition(nums []int, p, q int) int {
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

func randomPartition(nums []int, p, q int) int {
	i := rand.Intn(q-p+1) + p
	nums[i], nums[q] = nums[q], nums[i]
	return partition(nums, p, q)
}

func QuickSort(nums []int, p, q int) []int {
	if p >= q {
		return nums
	}
	pi := randomPartition(nums, p, q)
	QuickSort(nums, p, pi-1)
	QuickSort(nums, pi+1, q)
	return nums
}
