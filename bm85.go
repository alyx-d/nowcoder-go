package main

import (
	"strings"
)

func Bm85Solve(IP string) string {
	if isIPv4(IP) {
		return "IPv4"
	} else if isIPv6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLowerLetter(c byte) bool {
	return c >= 'a' && c <= 'f'
}

func isUpperLetter(c byte) bool {
	return c >= 'A' && c <= 'F'
}

func isIPv4(s string) bool {
	str := strings.Split(s, ".")
	if len(str) != 4 {
		return false
	}
	for i := 0; i < len(str); i++ {
		if len(str[i]) == 0 || len(str[i]) > 3 {
			return false
		}
		if len(str[i]) > 1 && str[i][0] == '0' {
			return false
		}
		num := 0
		for j := 0; j < len(str[i]); j++ {
			if !isNum(str[i][j]) {
				return false
			}
			num = num*10 + int(str[i][j]-'0')
		}
		if num > 255 {
			return false
		}
	}
	return true
}
func isIPv6(s string) bool {
	str := strings.Split(s, ":")
	if len(str) != 8 {
		return false
	}
	for i := 0; i < len(str); i++ {
		if len(str[i]) == 0 || len(str[i]) > 4 {
			return false
		}
		for j := 0; j < len(str[i]); j++ {
			if !isNum(str[i][j]) && !isLowerLetter(str[i][j]) && !isUpperLetter(str[i][j]) {
				return false
			}
		}
	}
	return true
}
