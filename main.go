package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(nums)
	i := BinarySearch(nums, 8)
	fmt.Println(i)
	fmt.Println(nums[i])
}
