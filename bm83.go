package main

func Trans(s string, n int) string {
	st := make([]string, 0)
	str := []byte(s)
	loru := func(c byte) byte {
		if c >= 'a' && c <= 'z' {
			c -= 32
			return c
		}
		if c >= 'A' && c <= 'Z' {
			c += 32
			return c
		}
		return c
	}
	for i := 0; i < n; i++ {
		str[i] = loru(str[i])
	}
	for i := 0; i < n; i++ {
		j := i
		for j < n && str[j] != ' ' {
			j++
		}
		st = append(st, string(str[i:j]))
		i = j
	}
	seq := ""
	ans := ""
	for i := len(st) - 1; i >= 0; i-- {
		ans += seq + st[i]
		seq = " "
	}
	if str[len(str)-1] == ' ' {
		return " " + ans
	}
	return ans
}
