package main

func RoateArr(n, m int, a []int) []int {
	reverse := func(l, r int) {
		for l < r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
	}
	m %= n
	reverse(0, n-1)
	reverse(0, m-1)
	reverse(m, n-1)
	return a
}
