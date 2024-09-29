package main

import "fmt"

func Bm86Solve(s, t string) string {
	reverse := func(str string) string {
		bytes := []byte(str)
		l, r := 0, len(str)-1
		for l < r {
			bytes[l], bytes[r] = bytes[r], bytes[l]
			l++
			r--
		}
		return string(bytes)
	}
	s = reverse(s)
	t = reverse(t)
	ans := ""
	carry := 0
	m, n := len(s), len(t)
	p, q := 0, 0
	for p < m || q < n || carry > 0 {
		sum := carry
		if p < m {
			sum += int(s[p] - '0')
			p++
		}
		if q < n {
			sum += int(t[q] - '0')
			q++
		}
		ans += fmt.Sprint(sum % 10)
		carry = sum / 10
	}
	return reverse(ans)
}
