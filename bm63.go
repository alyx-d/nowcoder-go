package main

func JumpFloor(number int) int {
	a, b := 1, 1
	for i := 0; i < number; i++ {
		a, b = b, a+b
	}
	return a
}
