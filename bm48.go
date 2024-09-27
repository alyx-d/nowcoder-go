package main

var data []int

func Insert(num int) {
	if len(data) == 0 {
		data = append(data, num)
	} else {
		idx := BinarySearch(data, num)
		data = append(data[:idx], append([]int{num}, data[idx:]...)...)
	}
}

func BinarySearch(nums []int, val int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < val {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r
}

func GetMedian() float64 {
	if len(data)&1 == 1 {
		return float64(data[len(data)>>1])
	} else {
		return float64((data[(len(data)-1)>>1] + data[len(data)>>1])) / 2
	}
}
