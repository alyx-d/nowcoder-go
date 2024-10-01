package main

func MaxLength(arr []int) int {
	dict := make(map[int]int, 0)
	ans := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for left, right := 0, 0; right < len(arr); right++ {
		if _, ok := dict[arr[right]]; !ok {
			dict[arr[right]] = 1
		} else {
			dict[arr[right]]++
		}
		for count := dict[arr[right]]; count > 1; count = dict[arr[right]] {
			dict[arr[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
