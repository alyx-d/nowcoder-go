package main

var st1 []int /* data */
var st2 []int /* min */

func Push2(val int) {
	st1 = append(st1, val)
	if len(st2) == 0 || st2[len(st2)-1] > val {
		st2 = append(st2, val)
	} else {
		st2 = append(st2, st2[len(st2)-1])
	}
}

func Pop2() {
	st1 = st1[:len(st1)-1]
	st2 = st2[:len(st2)-1]
}

func Top() int {
	return st1[len(st1)-1]
}

func Min() int {
	return st2[len(st2)-1]
}
