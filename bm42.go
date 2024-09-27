package main

var stack1 []int
var stack2 []int

func Push(val int) {
	stack1 = append(stack1, val)
}

func Pop() int {
	for len(stack2) > 0 {
		val := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		return val
	}
	for len(stack1) > 0 {
		stack2 = append(stack2, stack1[(len(stack1)-1)])
		stack1 = stack1[:len(stack1)-1]
	}
	val := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return val
}
