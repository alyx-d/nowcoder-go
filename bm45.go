package main

func MaxInWindow(num []int, size int) []int {
	ans := make([]int, 0)
	if len(num) == 0 || size == 0 || len(num) < size {
		return ans
	}
	deque := make([]int, 0)
	for i, val := range num {
		for len(deque) > 0 && val > num[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if deque[0]+size <= i {
			deque = deque[1:]
		}
		if i+1 >= size {
			ans = append(ans, num[deque[0]])
		}
	}
	return ans
}
