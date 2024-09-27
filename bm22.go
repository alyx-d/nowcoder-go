package main

import (
	"strconv"
	"strings"
)

func Compare(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	length := max(len(v1), len(v2))
	for i := 0; i < length; i++ {
		s1, s2 := 0, 0
		if i < len(v1) {
			num, _ := strconv.Atoi(v1[i])
			s1 += num
		}
		if i < len(v2) {
			num, _ := strconv.Atoi(v2[i])
			s2 += num
		}
		if s1 > s2 {
			return 1
		} else if s1 < s2 {
			return -1
		}
	}
	return 0
}
