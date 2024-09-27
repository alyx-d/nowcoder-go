package main

func FindKth(a []int, n, k int) int {
	l, r := 0, n-1
	for l <= r {
		pi := partition3(a, l, r)
		if pi+1 == k {
			return a[k-1]
		} else if pi+1 < k {
			l = pi + 1
		} else {
			r = pi - 1
		}
	}
	return 0
}

func partition3(nums []int, p, q int) int {
	pivot := nums[q]
	i := p - 1
	for j := p; j <= q-1; j++ {
		if nums[j] > pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[q] = nums[q], nums[i+1]
	return i + 1
}
