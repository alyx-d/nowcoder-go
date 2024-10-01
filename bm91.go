package main

func Bm91Reverse(s string) string {
	str := []byte(s)
	l, r := 0, len(str)-1
	for l < r {
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
	return string(str)
}
