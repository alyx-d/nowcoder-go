package main

func MoreThanHalfNum_Solution(numbers []int) int {
	count := 0
	candidate := 0
	for _, num := range numbers {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}
