package main

import "fmt"

type Pair struct {
	V1, V2 int
}

func m(pair Pair) {
	pair.V1++
	pair.V2--
	fmt.Println("inside", pair)
}

func main() {
	pair := Pair{1, 1}
	fmt.Println(pair)
	m(pair)
	fmt.Println(pair)
}
