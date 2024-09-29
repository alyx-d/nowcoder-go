package main

func GetLongestPalindrome(s string) int {
	expand := func(l, r int) int {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return r - l + 1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		ans = max(ans, max(expand(i, i), expand(i, i+1)))
	}
	return ans
}
